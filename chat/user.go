package chat

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

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
			fmt.Println("reading ...")
			fmt.Println(message)
			u.Global.Messages <- NewMessage(string(message), u.Username)
		}
	}
}

func (u *User) Write(message *Message) {
	b, _ := json.Marshal(message)

	if err := u.Conn.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Println("Error on write message:", err.Error())
	}
}
