package chat

import (
	"context"
	"encoding/json"
	"fluent/ent"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type User struct {
	Username string
	Conn     *websocket.Conn
	Room     *Chat
}

func (u *User) Read(client *ent.Client, userClient *ent.User) {
	for {
		if _, message, err := u.Conn.ReadMessage(); err != nil {
			log.Println("Error on read message: ", err.Error())
			break
		} else {
			fmt.Println("reading ...")
			fmt.Println(message)
			_, err := client.Message.
				Create().
				SetBody(string(message)).
				SetFrom(userClient).
				Save(context.Background())
			if err != nil {
				log.Print(err)
				break
			}

			u.Room.Messages <- NewMessage(string(message), u.Username)
		}
	}
}

func (u *User) Write(message *Message) {
	b, _ := json.Marshal(message)

	if err := u.Conn.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Println("Error on write message:", err.Error())
	}
}
