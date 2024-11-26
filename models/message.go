package models

type Message struct {
	ID       *int    `json:"id,omitempty"`
	Body     *string `json:"body,omitempty"`
	Sender   *string `json:"sender,omitempty"`
	Receiver *string `json:"receiver,omitempty"`
	ChatID   *int    `json:"chatId,omitempty"`
}
