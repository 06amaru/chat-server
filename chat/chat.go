package chat

import (
	"fmt"
)

type Chat struct {
	Users    map[string]*User
	Messages chan *Message
	Join     chan *User
	Leave    chan *User
	Id       string
}

func (c *Chat) Run() {
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
func (c *Chat) add(user *User) {
	if _, ok := c.Users[user.Username]; !ok {
		c.Users[user.Username] = user

		body := fmt.Sprintf("%s join the chat", user.Username)
		c.broadcast(NewMessage(body, "Server"))
	}
}

func (c *Chat) broadcast(message *Message) {
	fmt.Printf("Broadcast message: %v\n", message)
	for _, user := range c.Users {
		user.Write(message)
	}
}

func (c *Chat) disconnect(user *User) {
	if _, ok := c.Users[user.Username]; ok {
		defer user.Conn.Close()
		delete(c.Users, user.Username)

		body := fmt.Sprintf("%s left the chat", user.Username)
		c.broadcast(NewMessage(body, "Server"))
	}
}

func (c *Chat) GetMessages(chatID int) {

}
