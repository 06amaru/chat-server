package route

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/amaru0601/fluent/chat"
	"github.com/amaru0601/fluent/ent"
	"github.com/amaru0601/fluent/ent/user"
	"github.com/golang-jwt/jwt"
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

func (r *Route) JoinChat(manager map[string]*chat.Chat) echo.HandlerFunc {
	return func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		chatID := c.Param("id")
		if err != nil {
			log.Fatalln("Error on websocket connection:", err.Error())
		}
		defer ws.Close()

		//context has a map where user is the default key for auth-header
		authHeader := c.Get("user").(*jwt.Token)
		username := authHeader.Claims.(jwt.MapClaims)["username"].(string)
		userClient, err := r.db.User.
			Query().
			Where(user.UsernameEQ(username)).
			First(context.Background())
		if err != nil {
			log.Println(err)
		}

		// conseguir el chat
		if room, ok := manager[chatID]; ok {
			//conectar cliente con web socket
			user := &chat.User{
				Username: username,
				Conn:     ws,
				Room:     room,
			}

			room.Join <- user
			user.Read(r.db, userClient)
		} else {
			newRoom := &chat.Chat{
				Users:    make(map[string]*chat.User),
				Messages: make(chan *chat.Message),
				Join:     make(chan *chat.User),
				Leave:    make(chan *chat.User),
				Id:       chatID,
			}

			go newRoom.Run()

			manager[chatID] = newRoom

			user := &chat.User{
				Username: username,
				Conn:     ws,
				Room:     newRoom,
			}

			newRoom.Join <- user
			user.Read(r.db, userClient)
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
