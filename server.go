package main

import (
	"log"
	"net/http"

	"github.com/amaru0601/fluent/chat"
	"github.com/amaru0601/fluent/db"
	"github.com/amaru0601/fluent/route"
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

	var manager = make(map[string]*chat.Chat)

	// TODO: pass to authenticated routes with JWT
	e.GET("/chats/:id", r.JoinChat(manager))

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

		}
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
