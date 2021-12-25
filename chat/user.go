package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/amaru0601/fluent/ent"
	"github.com/gorilla/websocket"
)

type User struct {
	Username string
	Conn     *websocket.Conn
	Room     *Chat
}

func (u *User) Read(client *ent.Client, userClient *ent.User, chatID int) {
	defer func() {
		//necesitamos avisar al Chat que user se fue
		u.Room.Leave <- u
	}()
	for {
		if _, message, err := u.Conn.ReadMessage(); err != nil {
			log.Println("Error on read message =>\n", err.Error())
			break
		} else {
			fmt.Println("reading ...")
			fmt.Println(message)
			msg, err := client.Message.
				Create().
				SetBody(string(message)).
				SetFrom(userClient).
				Save(context.Background())
			if err != nil {
				log.Print(err)
				break
			}

			_, err = client.Chat.
				UpdateOneID(chatID).
				AddMessageIDs(msg.ID).
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
