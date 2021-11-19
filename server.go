package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/amaru0601/fluent/chat"
	"github.com/amaru0601/fluent/db"
	"github.com/amaru0601/fluent/route"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  512,
		WriteBufferSize: 512,
		CheckOrigin: func(r *http.Request) bool {
			log.Printf("%s %s%s %v\n", r.Method, r.Host, r.RequestURI, r.Proto)
			return r.Method == http.MethodGet
		},
	}
)

func main() {
	// Echo instance
	e := echo.New()

	// Ent client
	entClient, err := db.GetClient()
	if err != nil {
		log.Panicln("Database could not initialize")
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := route.NewRoute(entClient)

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	// encargado de almacenar las direcciones de memoria de los chats para luego poder conectarse
	var manager = make(map[string]*chat.Chat)

	//TODO HASH PASSWORD
	//curl -X POST -H 'Content-Type: application/json' -d '{"username":"jaoks", "password":"sdtc"}' localhost:1323/signup
	e.POST("/signup", r.SignUp())

	api := e.Group("/api")
	{
		auth := api.Group("/oauth")
		{
			// curl -X POST -H 'Content-Type: application/json' -d '{"username":"jaoks", "password":"sdtc"}' localhost:1323/api/oauth/signin
			auth.POST("/signin", r.SignIn())
		}
		fluent := api.Group("/fluent")
		{
			fluent.Use(middleware.JWT(route.MySigningKey))
			// wscat -c ws://localhost:1323/api/fluent/chat -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFtYXJ1IiwiZXhwIjoxNjM3Mzg5MTM2fQ.JysM4J-00sOP84Q_bzfW5wgw3QGPksSEikFe9JOVrAw"
			fluent.GET("/chat", r.JoinChat(manager))

			//TODO: Hacer endpoint para jalar todos los mensajes
		}
		plugged := api.Group("/plugged")
		{

			config := middleware.DefaultJWTConfig

			plugged.Use(func(next echo.HandlerFunc) echo.HandlerFunc {

				return func(c echo.Context) error {
					param := c.QueryParam("jwt")
					println(param)
					config.SigningKey = route.MySigningKey

					defaultKeyFunc := func(t *jwt.Token) (interface{}, error) {
						// Check the signing method
						if t.Method.Alg() != config.SigningMethod {
							return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
						}
						if len(config.SigningKeys) > 0 {
							if kid, ok := t.Header["kid"].(string); ok {
								if key, ok := config.SigningKeys[kid]; ok {
									return key, nil
								}
							}
							return nil, fmt.Errorf("unexpected jwt key id=%v", t.Header["kid"])
						}

						return config.SigningKey, nil
					}

					config.KeyFunc = defaultKeyFunc

					token, err := func(auth string, c echo.Context) (interface{}, error) {
						token_ := new(jwt.Token)
						var err error
						if _, ok := config.Claims.(jwt.MapClaims); ok {
							token_, err = jwt.Parse(auth, config.KeyFunc)
						} else {
							t := reflect.ValueOf(config.Claims).Type().Elem()
							claims := reflect.New(t).Interface().(jwt.Claims)
							token_, err = jwt.ParseWithClaims(auth, claims, config.KeyFunc)
						}
						if err != nil {
							return nil, err
						}
						if !token_.Valid {
							return nil, errors.New("invalid token")
						}
						return token_, nil
					}(param, c)

					if err == nil {
						// Store user information from token into context.
						c.Set(config.ContextKey, token)
						if config.SuccessHandler != nil {
							config.SuccessHandler(c)
						}
						return next(c)
					}
					if config.ErrorHandler != nil {
						return config.ErrorHandler(err)
					}
					if config.ErrorHandlerWithContext != nil {
						return config.ErrorHandlerWithContext(err, c)
					}
					return &echo.HTTPError{
						Code:     middleware.ErrJWTInvalid.Code,
						Message:  middleware.ErrJWTInvalid.Message,
						Internal: err,
					}
				}
			})
			plugged.GET("/chat", r.JoinChat(manager))
		}
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
