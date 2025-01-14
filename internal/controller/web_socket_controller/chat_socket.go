package websocketcontroller

import (
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/websocket/v2"
)

var (
	// Map to keep track of connected clients by `id`
	rooms = make(map[string][]*websocket.Conn)
	mu    sync.Mutex
)

// WebSocket endpoint that accepts `id` as a parameter

func SocketConnection(c *websocket.Conn) {
	fmt.Println("start socket")
	// Get the `id` parameter from the URL
	ChatKey, ok := c.Locals("ChatKey").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	id := ChatKey

	log.Printf("New connection with ID: %s", id)

	mu.Lock()
	// Check if the room already has 2 users
	if len(rooms[id]) >= 2 {
		// Reject the connection if the room is full
		mu.Unlock()
		err := c.WriteMessage(websocket.TextMessage, []byte("Room is full. Only 2 users allowed."))
		if err != nil {
			log.Println("Error sending rejection message:", err)
		}
		c.Close()
		return
	}

	// Add the WebSocket connection to the `rooms` map for the given `id`
	rooms[id] = append(rooms[id], c)
	mu.Unlock()

	// Ensure the client is removed from the room when the connection is closed
	defer func() {
		mu.Lock()
		// Remove the client from the room
		for i, conn := range rooms[id] {
			if conn == c {
				rooms[id] = append(rooms[id][:i], rooms[id][i+1:]...)
				break
			}
		}
		mu.Unlock()
	}()

	// Read and forward messages to other clients with the same `id`
	var (
		mt  int
		msg []byte
		err error
	)
	for {
		// Read incoming message
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("Received message: %s", msg)

		// Forward the message to other clients in the same `id` room
		mu.Lock()
		for _, client := range rooms[id] {
			// Don't send the message back to the sender
			if client != c {
				if err := client.WriteMessage(mt, msg); err != nil {
					log.Println("Error sending message to client:", err)
				}
			}
		}
		mu.Unlock()
	}
}
