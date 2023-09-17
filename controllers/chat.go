package controllers

import (
	"github.com/amaru0601/fluent/repository"
	"github.com/amaru0601/fluent/services"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  512,
		WriteBufferSize: 512,
		CheckOrigin: func(r *http.Request) bool {
			return r.Method == http.MethodGet
		},
	}
	// Get the JWT from the Authorization header
	//jwt := r.Header.Get("Authorization")

	// Validate the JWT
	// If the JWT is invalid, return an error
	//if !validateJWT(jwt) {
	//	return nil, fmt.Errorf("invalid JWT")
	//}

	return upgrader.Upgrade(w, r, nil)
}

type ChatController struct {
	svc services.ChatService
}

func NewChatController() ChatController {
	repo := repository.NewRepository()

	return ChatController{
		svc: services.NewChatService(repo),
	}
}

func (ctrl ChatController) GetMembers(c echo.Context) error {
	chatID, err := strconv.Atoi(c.QueryParam("chatID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	members, err := ctrl.svc.GetMembers(chatID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, members)
}

func (ctrl ChatController) GetChats(c echo.Context) error {
	//context has a map where "user" is the default key for jwt
	token := c.Get("user").(*jwt.Token)
	username := token.Claims.(jwt.MapClaims)["username"].(string)

	chats, err := ctrl.svc.GetChats(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, chats)
}

func (ctrl ChatController) CreateChat(c echo.Context) error {
	ws, err := Upgrade(c.Response(), c.Request())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer ws.Close()

	to := c.QueryParam("to")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	from := claims["username"].(string)

	err = ctrl.svc.CreateChat(to, from, ws)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return nil
}

func (ctrl ChatController) JoinChat(c echo.Context) error {
	ws, err := Upgrade(c.Response(), c.Request())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer ws.Close()

	chatID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	username := c.QueryParam("username")

	err = ctrl.svc.JoinChat(chatID, username, ws)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return nil
}
