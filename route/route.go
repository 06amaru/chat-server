package route

import (
	"fluent/ent"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

type Route struct {
	db *ent.Client
}

func NewRoute(client *ent.Client) *Route {
	return &Route{db: client}
}

func (r *Route) JoinChat() echo.HandlerFunc {
	return func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		chatID := c.Param("id")
		if err != nil {
			log.Fatalln("Error on websocket connection:", err.Error())
		}
		defer ws.Close()

		// conseguir el chat
		if room, ok := manager.chats[chatID]; ok {
			//conectar cliente con web socket
			//TODO: conseguir user desde JWT
			fmt.Println("chat saved ...")
			user := &User{
				Username: "jaoks",
				Conn:     ws,
				Global:   room,
			}

			room.Join <- user
			user.Read()

			//go chat.Run()
		} else {
			chat := &Chat{
				Users:    make(map[string]*User),
				Messages: make(chan *Message),
				Join:     make(chan *User),
				Leave:    make(chan *User),
			}

			//TODO: conseguir user desde JWT
			user := &User{
				Username: "amaru",
				Conn:     ws,
				Global:   chat,
			}
			manager.chats[chatID] = chat
			go chat.Run()

			fmt.Println("joining...")
			chat.Join <- user
			fmt.Println("joined user 1 ...")
			user.Read()
			fmt.Println("done ...")

		}

		return nil
	}
}
