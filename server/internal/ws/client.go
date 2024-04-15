package ws

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	// todo -> define client struct
	ID       string `json:"id"`
	Username string `json:"username"`
	Receiver string `json:"receiver"`
	Message  chan *Message
	Conn     *websocket.Conn
	RoomID   string `json:"roomId"`
}

type Message struct {
	// todo -> define msg struct
	Id        string    `json:"id"`
	Sender    string    `json:"sender"`
	Receiver  string    `json:"receiver"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
	RoomID    string    `json:"roomId"`
	Username  string    `json:"username"`
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(controller *Controller) {
	b := controller.board
	defer func() {
		b.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content:   string(m),
			Receiver:  c.Receiver,
			Sender:    c.Username,
			Timestamp: time.Now(),
			Id:        uuid.NewString(),
			RoomID:    c.RoomID,
			Username:  c.Username,
		}

		// todo: Put message to database
		query := "INSERT INTO messages(id, receiver,sender,content,created_at) VALUES ($1, $2,$3,$4,$5) returning id"
		var lastInsertedId string
		err = controller.db.QueryRowContext(context.Background(), query, msg.Id, msg.Receiver, msg.Sender, msg.Content, msg.Timestamp).Scan(&lastInsertedId)
		if err != nil {
			fmt.Printf("Failed to save message to database: %s", err)
			return
		}
		b.Broadcast <- msg
	}
}
