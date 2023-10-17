package services

import (
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
	return ChatService{repo: repo, chats: map[int]*apiChat.Chat{}}
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

func (svc ChatService) GetChats(username string) ([]*models.Chat, error) {
	chats, err := svc.repo.GetChats(username)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

type ChatModel struct {
	Chat     *ent.Chat
	Receiver *ent.User
	Sender   *ent.User
}

func (svc ChatService) VerifyChat(to, from string) (*ChatModel, error) {
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
		return nil, err
	}

	return &ChatModel{
		Chat:     chat,
		Receiver: receiver,
		Sender:   sender,
	}, nil
}

func (svc ChatService) CreateChat(sender, receiver *ent.User, ws *websocket.Conn) error {
	chat, err := svc.repo.CreateChat(sender.ID, receiver.ID)
	if err != nil {
		return err
	}

	svc.runChat(sender, ws, chat)

	return nil
}

func (svc ChatService) existChat(to, from string) (*ent.Chat, error) {
	chat, err := svc.repo.FindChatByUsernames(to, from)
	if chat != nil {
		return chat, nil
	}
	return nil, err
}

func (svc ChatService) JoinChat(chatID int, sender string, ws *websocket.Conn) error {
	user, err := svc.repo.FindUser(sender)
	if err != nil {
		return err
	}

	chat, err := svc.repo.FindChatByID(chatID, sender)
	if err != nil {
		return err
	}

	if room, ok := svc.chats[chatID]; ok {
		newUser := &apiChat.User{
			Username: sender,
			Conn:     ws,
			Room:     room,
		}
		room.Join <- newUser
		newUser.Read(svc.repo.Client, user, chatID)
	} else {
		svc.runChat(user, ws, chat)
	}

	return nil
}

func (svc ChatService) runChat(user *ent.User, ws *websocket.Conn, chat *ent.Chat) {
	newRoom := &apiChat.Chat{
		Users:    make(map[string]*apiChat.User),
		Messages: make(chan *apiChat.Message),
		Join:     make(chan *apiChat.User),
		Leave:    make(chan *apiChat.User),
		Id:       chat.ID,
	}

	svc.chats[chat.ID] = newRoom
	go newRoom.Run()

	newUser := &apiChat.User{
		Username: user.Username,
		Conn:     ws,
		Room:     newRoom,
	}

	newRoom.Join <- newUser
	newUser.Read(svc.repo.Client, user, chat.ID)
}
