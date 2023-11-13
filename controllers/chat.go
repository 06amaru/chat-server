package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amaru0601/fluent/ent"
	"github.com/amaru0601/fluent/repository"
	"github.com/amaru0601/fluent/services"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
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

func (ctrl ChatController) GetMessages(c echo.Context) error {
	chatID := c.Param("chatID")
	if chatID == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("there is no chat id found"))
	}

	chatIDInt, err := strconv.Atoi(chatID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("chat id is not a number"))
	}

	limit := c.Param("limit")
	if limit == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("limit required"))
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("limit is not a number"))
	}

	offset := c.Param("offset")
	if offset == "" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("offset required"))
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("offset is not a number"))
	}

	messages, err := ctrl.svc.GetMessages(chatIDInt, limitInt, offsetInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, messages)
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
	to := c.QueryParam("to")
	token := c.Get("user").(*jwt.Token)
	from := token.Claims.(jwt.MapClaims)["username"].(string)

	chat, err := ctrl.svc.VerifyChat(to, from)
	if chat != nil {
		return c.JSON(http.StatusBadRequest, chat)
	}
	switch err.(type) {
	case *ent.NotFoundError:
		fmt.Println("this chat can be created")
	default:
		return c.JSON(http.StatusInternalServerError, err)
	}

	ws, err := Upgrade(c.Response(), c.Request())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer ws.Close()
	err = ctrl.svc.CreateChat(chat.Receiver, chat.Sender, ws)
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

	token := c.Get("user").(*jwt.Token)
	from := token.Claims.(jwt.MapClaims)["username"].(string)

	err = ctrl.svc.JoinChat(chatID, from, ws)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return nil
}
