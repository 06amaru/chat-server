package main

import (
	"log"
	"net/http"
	"time"

	. "fluent/chat"
	"fluent/db"
	"fluent/route"

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

	var manager = make(map[string]*Chat)

	e.GET("/chats/:id", r.JoinChat(manager))
	e.POST("/login", login)

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

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

type jwtCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "jon" || password != "123456" {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		"username",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//TODO: put signing key to env variable
	t, err := token.SignedString([]byte("iosonic"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

type ChatManager struct {
	chats map[string]*Chat
}
