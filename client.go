package main

import "github.com/gorilla/websocket"

// client represents single chatting user
type client struct {
	// socket is web socket for this client
	socket *websocket.Conn
	// receive is a channel to receive messages from other clients
	receive chan []byte
	// room is the room this client chatting in
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
