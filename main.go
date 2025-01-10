package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/vishalpandhare01/holypeer_backend/internal"
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

	internal.SetUpRouts(app)

	// Start the server
	log.Fatal(app.Listen(":8000"))
}
