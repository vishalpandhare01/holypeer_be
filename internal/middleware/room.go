package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/holypeer_backend/initializer"
	"github.com/vishalpandhare01/holypeer_backend/internal/model"
)

func SecureRoom(C *fiber.Ctx) error {
	userId, ok := C.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	var member model.MemeberSchema
	if err := initializer.Db.Where("user_id = ?", userId).First(&member).Error; err != nil {
		fmt.Println("Errro.......", err.Error())
	}

	var listner model.ListenerSchema
	if err := initializer.Db.Where("user_id = ?", userId).First(&listner).Error; err != nil {
		fmt.Println("Errro.......", err.Error())
	}

	var chat model.Chat

	if err := initializer.Db.Where("member_id = ? OR listener_id = ?", member.ID, listner.ID).First(&chat).Error; err != nil {
		fmt.Println("Errro.......", err.Error())
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	C.Locals("ChatKey", chat.ChatKey)

	return C.Next()

}
