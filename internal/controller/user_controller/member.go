package usercontroller

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/holypeer_backend/initializer"
	"github.com/vishalpandhare01/holypeer_backend/internal/model"
	"github.com/vishalpandhare01/holypeer_backend/internal/utils/validation"
)

type AddMemberBody struct {
	UserID    string
	TodysFeel string
	Bio       string
	Country   string
}

// add member
func AddMemberProfile(C *fiber.Ctx) error {
	var body AddMemberBody
	userId, ok := C.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}
	body.UserID = userId
	fmt.Println(body)

	if err := C.BodyParser(&body); err != nil {
		return C.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if body.UserID == "" {
		return C.Status(400).JSON(fiber.Map{
			"message": "UserID required",
		})
	}

	var user model.User
	if err := initializer.Db.Where("id = ?", body.UserID).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return C.Status(404).JSON(fiber.Map{
				"message": "User: " + err.Error(),
			})
		}

		return C.Status(500).JSON(fiber.Map{
			"message": "User: " + err.Error(),
		})
	}

	var ExistMember model.MemeberSchema
	if err := initializer.Db.Where("user_id = ?", body.UserID).First(&ExistMember).Error; err != nil {
		if err.Error() != "record not found" {
			return C.Status(500).JSON(fiber.Map{
				"message": "Member: " + err.Error(),
			})
		}
	}
	if ExistMember.ID != "" {
		return C.Status(200).JSON(fiber.Map{
			"message": "Member: " + "record exist",
		})
	}

	if body.Bio == "" {
		return C.Status(400).JSON(fiber.Map{
			"message": "Bio required",
		})
	}

	if body.Country == "" {
		return C.Status(400).JSON(fiber.Map{
			"message": "Country required",
		})
	}

	if body.TodysFeel == "" {
		return C.Status(400).JSON(fiber.Map{
			"message": "TodysFeel required",
		})
	}

	if !validation.CheckTodysFeelExist(body.TodysFeel) {
		return C.Status(400).JSON(fiber.Map{
			"message": `"depression", "anxiety", "relationships", "ocd", "parenting", "family", "loneliness", "happy", "good" required`,
		})
	}

	var data = model.MemeberSchema{
		UserID:    body.UserID,
		TodysFeel: strings.ToLower(body.TodysFeel),
		Country:   body.Country,
		Bio:       body.Bio,
	}
	// strings.ToLower(text)
	if err := initializer.Db.Create(&data).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(201).JSON((fiber.Map{
		"message": "Your Member Profile Created SuccessfUlly",
		"data":    data,
	}))

}

// get member
func GetMemberById(C *fiber.Ctx) error {
	userId, ok := C.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	var member model.MemeberSchema
	if err := initializer.Db.
		Select("id", "user_id", "todys_feel", "bio", "country").
		Where("user_id", userId).
		Preload("User").
		First(&member).Error; err != nil {
		if err.Error() == "record not found" {
			return C.Status(404).JSON(fiber.Map{
				"message": "member: " + err.Error(),
			})
		}
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    member,
	})
}
