package route

import (
	"context"
	. "fluent/chat"
	"fluent/ent"
	"fluent/ent/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
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

type Route struct {
	db *ent.Client
}

func NewRoute(client *ent.Client) *Route {
	return &Route{db: client}
}

// originalmente el chat id puede ser null o int
// si el chat id es null es porque el cliente esta empezando una nueva conversacion
// se debe crear un historial del chat en la base de datos
// si el chat id es int conseguir el historial y conseguir el chat desde el manager

func (r *Route) JoinChat(manager map[string]*Chat) echo.HandlerFunc {
	return func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		chatID := c.Param("id")
		if err != nil {
			log.Fatalln("Error on websocket connection:", err.Error())
		}
		defer ws.Close()

		//TODO: conseguir user desde JWT
		userClient, err := r.db.User.
			Query().
			Where(user.UsernameEQ("jaoks")).
			First(context.Background())
		if err != nil {
			fmt.Println(err)
		}

		// conseguir el chat
		if room, ok := manager[chatID]; ok {
			//conectar cliente con web socket
			//TODO: conseguir user desde JWT
			fmt.Println("chat saved ...")
			user := &User{
				Username: "jaoks",
				Conn:     ws,
				Room:     room,
			}

			room.Join <- user
			user.Read(r.db, userClient)

			//go chat.Run()
		} else {
			chat := &Chat{
				Users:    make(map[string]*User),
				Messages: make(chan *Message),
				Join:     make(chan *User),
				Leave:    make(chan *User),
				Id:       chatID,
			}

			user := &User{
				Username: "amaru",
				Conn:     ws,
				Room:     chat,
			}
			manager[chatID] = chat

			fmt.Println("Crear chat ...")
			// crear chat en la bd
			if chatClient, err := r.db.Chat.
				Create().
				SetName(chatID).
				SetType("private").
				SetDeleted(false).
				Save(context.Background()); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(chatClient)
			}

			/*
				if err != nil {
					return c.String(http.StatusBadRequest, err.Error())
				}*/

			go chat.Run()

			fmt.Println("joining...")
			chat.Join <- user
			fmt.Println("joined user 1 ...")
			user.Read(r.db, userClient)
			fmt.Println("done ...")

		}

		return nil
	}
}
