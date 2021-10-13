package main

import (
	"log"
	"net/http"

	. "fluent/chat"
	"fluent/db"
	"fluent/route"

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

	var manager = make(map[string]*Chat)

	e.GET("/chats/:id", r.JoinChat(manager))

	/*r := e.Group("/auth")
	{
		config := middleware.JWTConfig{
			Claims:     &jwtCustomClaims{},
			SigningKey: []byte("iosonic"),
		}
		r.Use(middleware.JWTWithConfig(config))
		r.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!\n")
		})
	}*/

	api := e.Group("/api")
	{
		auth := api.Group("/oauth")
		{
			// TODO finish sign in
			auth.POST("/signin", r.SignIn())
			// TODO return JWT
		}
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
