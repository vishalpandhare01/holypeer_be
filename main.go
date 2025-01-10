package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/joho/godotenv"
	"github.com/vishalpandhare01/holypeer_backend/initializer"
	"github.com/vishalpandhare01/holypeer_backend/internal"
	"github.com/vishalpandhare01/holypeer_backend/internal/migration"
)

func Init() {
	fmt.Println("Db connnection")
	initializer.ConnectDatabase()
	migration.DbMigration()
}

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Env not loaded: ", err)
	}
	Init()
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
