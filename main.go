package main

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var (
	// Map to keep track of connected clients by `id`
	rooms = make(map[string][]*websocket.Conn)
	mu    sync.Mutex
)

func main() {
	app := fiber.New()

	// WebSocket middleware for upgrading the connection
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// WebSocket endpoint that accepts `id` as a parameter
	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// Get the `id` parameter from the URL
		id := c.Params("id")
		log.Printf("New connection with ID: %s", id)

		// Add the WebSocket connection to the `rooms` map for the given `id`
		mu.Lock()
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
	}))

	// Start the server
	log.Fatal(app.Listen(":8000"))
}
