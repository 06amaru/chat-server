package services

import (
	"fmt"
	apiChat "github.com/amaru0601/fluent/chat"
	"github.com/amaru0601/fluent/ent"
	"github.com/amaru0601/fluent/models"
	"github.com/amaru0601/fluent/repository"
	"github.com/gorilla/websocket"
)

type ChatService struct {
	repo  *repository.Repository
	chats map[int]*apiChat.Chat
}

func NewChatService(repo *repository.Repository) ChatService {
	return ChatService{repo: repo}
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

func (svc ChatService) GetChats(username string) ([]*ent.Chat, error) {
	chats, err := svc.repo.GetChats(username)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (svc ChatService) CreateChat(to, from string, ws *websocket.Conn) error {
	receiver, err := svc.repo.FindUser(to)
	if err != nil {
		return err
	}

	sender, err := svc.repo.FindUser(from)
	if err != nil {
		return err
	}

	err = svc.existChat(to, from)
	if err != nil {
		return err
	}

	chat, err := svc.repo.CreateChat(receiver.ID, sender.ID)
	if err != nil {
		return err
	}

	newRoom := &apiChat.Chat{
		Users:    make(map[string]*apiChat.User),
		Messages: make(chan *apiChat.Message),
		Join:     make(chan *apiChat.User),
		Leave:    make(chan *apiChat.User),
		Id:       chat.ID,
	}

	svc.chats[chat.ID] = newRoom
	go newRoom.Run()

	user := &apiChat.User{
		Username: from,
		Conn:     ws,
		Room:     newRoom,
	}

	newRoom.Join <- user
	user.Read(svc.repo.Client, sender, chat.ID)

	return nil
}

func (svc ChatService) existChat(to, from string) error {
	chat, err := svc.repo.FindChatByUsernames(to, from)
	if chat != nil {
		return fmt.Errorf("chat already exists")
	}
	switch err.(type) {
	case *ent.NotFoundError:
		return nil
	default:
		return err
	}
}

func (svc ChatService) JoinChat(chatID int, username string, ws *websocket.Conn) error {
	user, err := svc.repo.FindUser(username)
	if err != nil {
		return err
	}

	if room, ok := svc.chats[chatID]; ok {
		userChat := &apiChat.User{
			Username: username,
			Conn:     ws,
			Room:     room,
		}
		room.Join <- userChat
		userChat.Read(svc.repo.Client, user, chatID)
		return nil
	} else {
		return fmt.Errorf("chat not found")
	}
}
