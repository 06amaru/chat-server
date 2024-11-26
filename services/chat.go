package services

import (
	"errors"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/jaox1/chat-server/chat"
	"github.com/jaox1/chat-server/ent"
	"github.com/jaox1/chat-server/models"
	"github.com/jaox1/chat-server/repository"
)

type ChatService struct {
	repo          *repository.Repository
	socketManager *chat.SocketManager
}

func NewChatService(repo *repository.Repository) ChatService {
	newSocketManager := chat.NewSocketManager()
	go newSocketManager.Run()
	return ChatService{repo: repo, socketManager: newSocketManager}
}

func (svc ChatService) GetMembers(chatID int) ([]models.User, error) {
	members, err := svc.repo.GetChatMembers(chatID)
	if err != nil {
		return nil, err
	}

	membersArr := make([]models.User, 0)
	for _, v := range members {
		var m models.User
		m.Username = v.Username
		m.PublicKey = v.PublicKey
		membersArr = append(membersArr, m)
	}

	return membersArr, nil
}

func (svc ChatService) GetMessages(chatID, limit, offset int) (interface{}, error) {
	messages, err := svc.repo.GetMessages(chatID, limit, offset)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (svc ChatService) GetChats(username string) ([]*models.Chat, error) {
	chats, err := svc.repo.GetChats(username)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (svc ChatService) VerifyChat(to, from string) (*models.EntChat, error) {
	receiver, err := svc.repo.FindUser(to)
	if err != nil {
		return nil, err
	}

	sender, err := svc.repo.FindUser(from)
	if err != nil {
		return nil, err
	}

	chat, err := svc.existChat(to, from)
	if err != nil {
		var notFoundError *ent.NotFoundError
		switch {
		case errors.As(err, &notFoundError):
			// this chat can be created
			return &models.EntChat{
				Chat:     chat,
				Receiver: receiver,
				Sender:   sender,
			}, nil
		default:
			return nil, err
		}
	}

	return nil, fmt.Errorf("chat with ID %d exists between %s and %s", chat.ID, to, from)
}

func (svc ChatService) CreateChat(sender, receiver *ent.User) error {
	_, err := svc.repo.CreateChat(sender.ID, receiver.ID)
	if err != nil {
		return err
	}

	return nil
}

func (svc ChatService) existChat(to, from string) (*ent.Chat, error) {
	chat, err := svc.repo.FindChatByUsernames(to, from)
	if chat != nil {
		return chat, nil
	}
	return nil, err
}

func (svc ChatService) Subscribe(username string, ws *websocket.Conn) error {
	user, err := svc.repo.FindUser(username)
	if err != nil {
		return err
	}

	newUser := &chat.User{
		Username:      username,
		Conn:          ws,
		SocketManager: svc.socketManager,
		Database:      svc.repo.Client,
		EntUser:       user,
	}
	svc.socketManager.Join <- newUser

	return newUser.Listen()
}
