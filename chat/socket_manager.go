package chat

import (
	"fmt"

	"github.com/jaox1/chat-server/models"
)

/*
The socket manager stores all user connections. Each incoming message
will be forwarded by the socket manager to the corresponding user.
*/
type SocketManager struct {
	Messages chan *models.Message
	Join     chan *User
	Leave    chan *User
	Id       int
	Users    map[string]*User
}

func NewSocketManager() *SocketManager {
	return &SocketManager{
		Messages: make(chan *models.Message),
		Join:     make(chan *User),
		Leave:    make(chan *User),
		Users:    make(map[string]*User),
	}
}

func (c *SocketManager) Run() {
	fmt.Println("running chat ... ")
	for {
		select {
		case user := <-c.Join:
			c.add(user)
		case message := <-c.Messages:
			c.broadcast(message)
		case user := <-c.Leave:
			c.disconnect(user)
		}
	}
}
func (c *SocketManager) add(user *User) {
	if _, ok := c.Users[user.Username]; !ok {
		c.Users[user.Username] = user

		body := fmt.Sprintf("%s is online", user.Username)
		sender := user.Username
		c.broadcast(&models.Message{
			Body:   &body,
			Sender: &sender,
		})
	}
}

func (c *SocketManager) broadcast(message *models.Message) {
	if message.Receiver == nil {
		// offline and online notification to all user
		for _, user := range c.Users {
			user.Send(message)
		}
		return
	}

	if user, ok := c.Users[*message.Sender]; ok {
		user.Send(message)
	}

	if user, ok := c.Users[*message.Receiver]; ok {
		user.Send(message)
	}
}

func (c *SocketManager) disconnect(user *User) {
	if _, ok := c.Users[user.Username]; ok {
		defer user.Conn.Close()
		delete(c.Users, user.Username)

		body := fmt.Sprintf("%s is offline", user.Username)
		sender := user.Username
		c.broadcast(&models.Message{
			Body:   &body,
			Sender: &sender,
		})
	}
}
