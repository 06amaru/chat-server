package main

import (
	"context"
	"log"
	"net/http"

	"github.com/amaru0601/fluent/chat"
	"github.com/amaru0601/fluent/db"
	"github.com/amaru0601/fluent/ent/user"
	"github.com/amaru0601/fluent/route"
	"github.com/amaru0601/fluent/security"
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

type DataStruct struct {
	Ar []uint8 `json:"data"`
}

type PrivateKey struct {
	Pk DataStruct `json:"privateKey"`
}

func main() {
	// Echo instance
	e := echo.New()
	e.Use(middleware.CORS())
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
	var manager = make(map[int]*chat.Chat)

	// encargado de almacenar las llaves privadas de los usuarios
	var keeper = make(map[string][]uint8)

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

			fluent.GET("/username", func(c echo.Context) error {
				authHeader := c.Get("user").(*jwt.Token)
				username := authHeader.Claims.(jwt.MapClaims)["username"].(string)
				return c.JSON(http.StatusAccepted, username)
			})

			fluent.GET("/public-key", func(c echo.Context) error {
				username := c.QueryParam("username")
				user, _ := entClient.User.Query().Where(
					user.UsernameEQ(username),
				).First(context.Background())
				return c.JSON(http.StatusOK, user.PublicKey)
			})

			fluent.GET("/secret-key", func(c echo.Context) error {

				authHeader := c.Get("user").(*jwt.Token)
				username := authHeader.Claims.(jwt.MapClaims)["username"].(string)
				user, _ := entClient.User.Query().Where(
					user.UsernameEQ(username),
				).First(context.Background())
				return c.JSON(http.StatusAccepted, user.PrivateKey)
			})

			fluent.POST("/secret-key", func(c echo.Context) error {
				k := new(PrivateKey)
				if err := c.Bind(&k); err != nil {
					log.Println("ERROR ")
					log.Println(err)
				}

				//context has a map where user is the default key for auth-header
				authHeader := c.Get("user").(*jwt.Token)
				username := authHeader.Claims.(jwt.MapClaims)["username"].(string)
				keeper[username] = k.Pk.Ar
				return c.String(http.StatusOK, "ok")
			})

			//consigue todos los chats
			fluent.GET("/chats", func(c echo.Context) error {
				//context has a map where user is the default key for auth-header
				authHeader := c.Get("user").(*jwt.Token)
				username := authHeader.Claims.(jwt.MapClaims)["username"].(string)
				entClient, err := db.GetClient()

				if err != nil {
					return c.String(http.StatusInternalServerError, "no se pudo conectar a la base de datos")
				}

				user, _ := entClient.User.Query().Where(
					user.UsernameEQ(username),
				).First(context.Background())

				chats, _ := user.QueryChats().All(context.Background())

				return c.JSON(http.StatusOK, chats)
			})

			//TODO: Hacer endpoint para jalar todos los mensajes -- CHALLENGE: almacenar las KEYS de esos mensajes
		}
		plugged := api.Group("/plugged")
		{
			plugged.Use(security.CustomMiddleware)
			plugged.GET("/chat", r.JoinChat(manager))
		}
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
