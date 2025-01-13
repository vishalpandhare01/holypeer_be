package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func IsMember(C *fiber.Ctx) error {
	role, ok := C.Locals("userType").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userType is not a string")
	}

	if role == "member" {
		return C.Next()
	}
	return C.Status(401).JSON(fiber.Map{
		"message": "You Are Not Authorized for this operation please login as member",
	})
}

func IsListener(C *fiber.Ctx) error {
	role, ok := C.Locals("userType").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userType is not a string")
	}

	if role == "listner" {
		return C.Next()
	}
	return C.Status(401).JSON(fiber.Map{
		"message": "You Are Not Authorized for this operation please login as listner",
	})
}
