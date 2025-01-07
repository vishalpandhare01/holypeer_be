package internal

import (
	"github.com/gofiber/fiber/v2"
	usercontroller "github.com/vishalpandhare01/holypeer_backend/internal/user_controller"
)

func SetUpRouts(app *fiber.App) {
	app.Get("/", usercontroller.Server)
}
