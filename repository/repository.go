package repository

import (
	"context"

	"github.com/jaox1/chat-server/db"
	"github.com/jaox1/chat-server/ent"
	chatEnt "github.com/jaox1/chat-server/ent/chat"
	"github.com/jaox1/chat-server/ent/message"
	userEnt "github.com/jaox1/chat-server/ent/user"
	"github.com/jaox1/chat-server/models"
)

type Repository struct {
	Client *ent.Client
}

func NewRepository() *Repository {
	postgres := db.GetPostgresClient()

	return &Repository{
		Client: postgres,
	}
}

func (repo Repository) SignUp(username, password string, privateK, publicK []byte) error {
	_, err := repo.Client.User.
		Create().
		SetUsername(username).
		SetPassword(password).
		SetPrivateKey(privateK).
		SetPublicKey(publicK).
		Save(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func (repo Repository) FindUser(username string) (*ent.User, error) {
	user, err := repo.Client.User.
		Query().
		Where(userEnt.UsernameEQ(username)).
		First(context.Background())

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo Repository) GetChatMembers(chatID int) ([]*ent.User, error) {
	chat, err := repo.Client.Chat.
		Get(context.Background(), chatID)
	if err != nil {
		return nil, err
	}

	members, err := chat.QueryMembers().
		Select(userEnt.FieldUsername, userEnt.FieldPublicKey).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	return members, nil
}

func (repo Repository) GetChats(username string) ([]*models.Chat, error) {
	user, err := repo.Client.User.Query().Where(userEnt.UsernameEQ(username)).
		WithChats(func(query *ent.ChatQuery) {
			query.WithMembers()
		}).First(context.Background())
	if err != nil {
		return nil, err
	}

	var response []*models.Chat
	for _, chat := range user.Edges.Chats {
		for _, member := range chat.Edges.Members {
			if member.Username != username {
				newChat := &models.Chat{
					ChatID: chat.ID,
					Sender: member.Username,
				}
				response = append(response, newChat)
			}
		}
	}

	return response, nil
}

func (repo Repository) FindChatByUsernames(to, from string) (*ent.Chat, error) {
	chat, err := repo.Client.Chat.
		Query().
		Where(chatEnt.And(
			chatEnt.HasMembersWith(userEnt.UsernameEQ(to)),
			chatEnt.HasMembersWith(userEnt.UsernameEQ(from)),
		)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (repo Repository) FindChatByID(chatID int, from string) (*ent.Chat, error) {
	chat, err := repo.Client.Chat.
		Query().
		Where(chatEnt.And(
			chatEnt.HasMembersWith(userEnt.UsernameEQ(from)),
			chatEnt.ID(chatID),
		)).
		Only(context.Background())

	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (repo Repository) CreateChat(to, from int) (*ent.Chat, error) {
	chat, err := repo.Client.Chat.Create().
		SetType("public").
		AddMemberIDs(to, from).Save(context.Background())

	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (repo Repository) GetMessages(chatID, limit, offset int) ([]*ent.Message, error) {
	chat, err := repo.Client.Chat.Query().Where(chatEnt.ID(chatID)).Only(context.Background())
	if err != nil {
		return nil, err
	}
	messages, err := repo.Client.Chat.QueryMessages(chat).
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(message.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return messages, nil
}
