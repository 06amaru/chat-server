package main

import (
	"log"
	"net/http"

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

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/chat/:chatId", ChatManager)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

type User struct {
	Username string
	Conn     *websocket.Conn
	Global   *Chat
}

type Message struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	Sender string `json:"sender"`
}

type Chat struct {
	users   map[string]*User
	message chan *Message
	join    chan *User
	leave   chan *User
}

func (chat *Chat) Handler(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatalln("Error on websocket connection:", err.Error())
		return err
	}
	// keys := c.Req
	defer ws.Close()
	// for {
	// 	err := ws.
	// }
}

func ChatManager(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	//TODO
}
