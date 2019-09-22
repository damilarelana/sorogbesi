package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Client defines template of the server-end representation of each Client connection
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// Message defines template for that message structure being passed between Clients
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {

	// Unregister the client once mesage is passed
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// extract mesage from connection
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// construct the message
		message := Message{
			Type: messageType,
			Body: string(p),
		}

		// pass the message to other Clients
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
