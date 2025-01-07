package usercontroller

import "github.com/gofiber/fiber/v2"

func Server(C *fiber.Ctx) error {
	return C.Status(200).JSON(fiber.Map{
		"message": "server running now",
	})
}
