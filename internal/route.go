package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	usercontroller "github.com/vishalpandhare01/holypeer_backend/internal/user_controller"
	websocketcontroller "github.com/vishalpandhare01/holypeer_backend/internal/web_socket_controller"
)

func SetUpRouts(app *fiber.App) {
	app.Get("/", usercontroller.Server)
	app.Get("/ws/:id", websocket.New(websocketcontroller.SocketConnection))
}
