package route

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/amaru0601/fluent/chat"
	"github.com/amaru0601/fluent/ent"
	"github.com/amaru0601/fluent/ent/user"
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

func (r *Route) JoinChat(manager map[string]*chat.Chat) echo.HandlerFunc {
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
			user := &chat.User{
				Username: "jaoks",
				Conn:     ws,
				Room:     room,
			}

			room.Join <- user
			user.Read(r.db, userClient)

			//go chat.Run()
		} else {
			newChat := &chat.Chat{
				Users:    make(map[string]*chat.User),
				Messages: make(chan *chat.Message),
				Join:     make(chan *chat.User),
				Leave:    make(chan *chat.User),
				Id:       chatID,
			}

			user := &chat.User{
				Username: "amaru",
				Conn:     ws,
				Room:     newChat,
			}
			manager[chatID] = newChat

			fmt.Println("Crear chat ...")
			// crear chat en la bd
			chatClient, err := r.db.Chat.
				Create().
				SetName(chatID).
				SetType("private").
				SetDeleted(false).
				Save(context.Background())
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(chatClient)
			}

			/*
				if err != nil {
					return c.String(http.StatusBadRequest, err.Error())
				}*/

			go newChat.Run()

			fmt.Println("joining...")
			newChat.Join <- user
			fmt.Println("joined user 1 ...")
			user.Read(r.db, userClient)
			fmt.Println("done ...")

		}

		return nil
	}
}

type UserSignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Route) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(UserSignIn)
		if err := c.Bind(u); err != nil {
			return err
		}

		userEnt, err := r.db.User.
			Query().
			Where(user.UsernameEQ(u.Username)).
			First(context.Background())

		//TODO : mejorar mensaje de error por parte de ent user not found
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		if u.Password != userEnt.Password {
			return c.String(http.StatusUnauthorized, "wrong password")
		}

		jwtoken, err := makeToken(u)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		return c.String(http.StatusOK, jwtoken)
	}
}

func (r *Route) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(UserSignIn)
		if err := c.Bind(u); err != nil {
			return err
		}

		_, err := r.db.User.
			Create().
			SetUsername(u.Username).
			SetPassword(u.Password).
			Save(context.Background())
		if err != nil {
			fmt.Println(err)
			return err
		}

		return c.String(http.StatusOK, "user has been created")
	}
}
