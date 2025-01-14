package chatcontroller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vishalpandhare01/holypeer_backend/initializer"
	"github.com/vishalpandhare01/holypeer_backend/internal/model"
)

func RequestForchat(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	var member model.MemeberSchema

	if err := initializer.Db.Where("user_id = ?", userId).First(&member).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data := model.Chat{
		MemberID: member.ID,
		ChatKey:  uuid.New().String(),
	}

	if err := initializer.Db.Create(&data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "success",
		"data":    data,
	})
}

func GetCahtRequests(c *fiber.Ctx) error {
	var chat []model.Chat

	if err := initializer.Db.
		Preload("Member").
		Preload("Member.User").
		Preload("Listener.User").
		Select("id", "member_id", "created_at"). // Select only member_id and listener_id
		Find(&chat).
		Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "succss",
		"data":    chat,
	})
}

func AcceptRequest(c *fiber.Ctx) error {
	chatId := c.Params("chatId")

	var chat model.Chat
	if err := initializer.Db.Where("id = ?", chatId).First(&chat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if chat.IsAccepted {
		return c.Status(400).JSON(fiber.Map{
			"message": "Chat request already accepted",
		})
	}

	userId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	var listner model.ListenerSchema

	if err := initializer.Db.Where("user_id = ?", userId).First(&listner).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	chat.ListenerID = &listner.ID
	chat.IsAccepted = true

	if err := initializer.Db.Save(&chat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "success",
	})
}

func CloseChat(c *fiber.Ctx) error {
	chatId := c.Params("id")
	userId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	var member model.Chat
	if err := initializer.Db.Where("user_id = ?", userId).First(&member).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var chat model.Chat
	if err := initializer.Db.Where("id = ? and member_id = ? ", chatId, member.ID).Delete(&chat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "chat closed succssfully",
	})
}
