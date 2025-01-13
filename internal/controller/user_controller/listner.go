package usercontroller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/holypeer_backend/initializer"
	"github.com/vishalpandhare01/holypeer_backend/internal/model"
)

type AddListnerBody struct {
	UserID    string
	TodysFeel string
	Bio       string
	Country   string
}

// add Listner
func AddListnerProfile(C *fiber.Ctx) error {
	var body AddListnerBody
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

	var ExistListner model.ListenerSchema
	if err := initializer.Db.Where("user_id = ?", body.UserID).First(&ExistListner).Error; err != nil {
		if err.Error() != "record not found" {
			return C.Status(500).JSON(fiber.Map{
				"message": "Listner: " + err.Error(),
			})
		}
	}
	if ExistListner.ID != "" {
		return C.Status(200).JSON(fiber.Map{
			"message": "Listner: " + "record exist",
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

	var data = model.ListenerSchema{
		UserID:  body.UserID,
		Country: body.Country,
		Bio:     body.Bio,
	}

	// strings.ToLower(text)
	if err := initializer.Db.Create(&data).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(201).JSON((fiber.Map{
		"message": "Your Listner Profile Created SuccessfUlly",
		"data":    data,
	}))

}

// get Listner
func GetListnerById(C *fiber.Ctx) error {
	userId, ok := C.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	var Listner model.ListenerSchema
	if err := initializer.Db.
		Select("id", "user_id", "bio", "country").
		Where("user_id", userId).
		Preload("User").
		First(&Listner).Error; err != nil {
		if err.Error() == "record not found" {
			return C.Status(404).JSON(fiber.Map{
				"message": "Listner: " + err.Error(),
			})
		}
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    Listner,
	})
}
