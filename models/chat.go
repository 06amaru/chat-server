package models

import "github.com/jaox1/chat-server/ent"

type Chat struct {
	ChatID   int    `json:"chatID"`
	Username string `json:"username"`
}

type EntChat struct {
	Chat     *ent.Chat
	Receiver *ent.User
	Sender   *ent.User
}
