package main

import (
	"encoding/json"
	"fmt"
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

	chatManager := &ChatManager{
		chats: make(map[string]*Chat),
	}

	e.GET("/chats/:id", chatManager.handle)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

type User struct {
	Username string
	Conn     *websocket.Conn
	Global   *Chat
}

func (u *User) Read() {
	for {
		if _, message, err := u.Conn.ReadMessage(); err != nil {
			log.Println("Error on read message: ", err.Error())
			break
		} else {
			u.Global.messages <- NewMessage(string(message), u.Username)
		}
	}
}

func (u *User) Write(message *Message) {
	b, _ := json.Marshal(message)

	if err := u.Conn.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Println("Error on write message:", err.Error())
	}
}

type Message struct {
	ID     int64  `json:"id"`
	Body   string `json:"body"`
	Sender string `json:"sender"`
}

func NewMessage(body string, sender string) *Message {
	return &Message{
		ID:     1,
		Body:   body,
		Sender: sender,
	}
}

type Chat struct {
	users    map[string]*User
	messages chan *Message
	join     chan *User
	leave    chan *User
}

func (c *Chat) Run() {
	for {
		select {
		case user := <-c.join:
			c.add(user)
		case message := <-c.messages:
			c.broadcast(message)
		case user := <-c.leave:
			c.disconnect(user)
		}
	}
}
func (c *Chat) add(user *User) {
	if _, ok := c.users[user.Username]; !ok {
		c.users[user.Username] = user

		body := fmt.Sprintf("%s join the chat", user.Username)
		c.broadcast(NewMessage(body, "Server"))
	}
}

func (c *Chat) broadcast(message *Message) {
	log.Printf("Broadcast message: %v\n", message)
	for _, user := range c.users {
		user.Write(message)
	}
}

func (c *Chat) disconnect(user *User) {
	if _, ok := c.users[user.Username]; ok {
		defer user.Conn.Close()
		delete(c.users, user.Username)

		body := fmt.Sprintf("%s left the chat", user.Username)
		c.broadcast(NewMessage(body, "Server"))
	}
}

type ChatManager struct {
	chats map[string]*Chat
}

func (manager *ChatManager) handle(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	chatID := c.Param("id")
	if err != nil {
		log.Fatalln("Error on websocket connection:", err.Error())
	}
	defer ws.Close()

	// conseguir el chat
	if chat, ok := manager.chats[chatID]; ok {
		//conectar cliente con web socket
		//TODO: conseguir user desde JWT
		user := &User{
			Username: "jaoks",
			Conn:     ws,
			Global:   chat,
		}

		chat.join <- user
		user.Read()

		go chat.Run()
	} else {
		chat := &Chat{
			users:    make(map[string]*User),
			messages: make(chan *Message),
			join:     make(chan *User),
			leave:    make(chan *User),
		}

		//TODO: conseguir user desde JWT
		user := &User{
			Username: "amaru",
			Conn:     ws,
			Global:   chat,
		}

		chat.join <- user
		user.Read()

		go chat.Run()
	}

	return nil
}
