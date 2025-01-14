package usercontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/holypeer_backend/initializer"
	"github.com/vishalpandhare01/holypeer_backend/internal/model"
	"github.com/vishalpandhare01/holypeer_backend/internal/utils"
	"github.com/vishalpandhare01/holypeer_backend/internal/utils/jwtToken"
	"github.com/vishalpandhare01/holypeer_backend/internal/utils/validation"
)

func Server(C *fiber.Ctx) error {
	return C.Status(200).JSON(fiber.Map{
		"message": "server running now",
	})
}

type UserBody struct {
	Name        string
	Email       string
	DateOfBirth string
	IsListner   bool
}

func RegisterUser(C *fiber.Ctx) error {
	var body UserBody

	if err := C.BodyParser(&body); err != nil {
		return C.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if body.Name == "" {
		return C.Status(400).JSON(fiber.Map{
			"message": "Name is required",
		})
	}

	if body.Email == "" {
		return C.Status(400).JSON(fiber.Map{
			"message": "Email is required",
		})
	}

	if body.DateOfBirth == "" {
		return C.Status(400).JSON(fiber.Map{
			"message": "DOB is required",
		})
	}

	if !validation.ValidDateOfBirth(body.DateOfBirth) {
		return C.Status(400).JSON(fiber.Map{
			"message": "YYYY-MM-DD formate is required",
		})
	}

	if validation.CheckNameExist(body.Name) {
		return C.Status(400).JSON(fiber.Map{
			"message": "Use Another New Name",
		})
	}

	if validation.CheckEmailExist(body.Email) {
		return C.Status(400).JSON(fiber.Map{
			"message": "You already have acount with this email",
		})
	}

	var data = model.User{
		Name:        body.Name,
		Email:       body.Email,
		DateOfBirth: body.DateOfBirth,
		IsListner:   body.IsListner,
	}

	if err := initializer.Db.Create(&data).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(201).JSON(fiber.Map{
		"message": "success",
		"data":    data,
	})

}

// send otp for authentication
type sendOtpBody struct {
	Email string
}

func SendOtp(C *fiber.Ctx) error {
	var body sendOtpBody

	if err := C.BodyParser(&body); err != nil {
		return C.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var user model.User

	if err := initializer.Db.Where("email = ?", body.Email).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return C.Status(404).JSON(fiber.Map{
				"message": "Email: " + err.Error(),
			})
		}
		return C.Status(500).JSON(fiber.Map{
			"message": "Email: " + err.Error(),
		})
	}

	code := utils.Otp_Number_Generate()
	//Todo :- sent otp to email not in response
	var userOtp model.User_Otp
	//check if user send otp already then create new else update existing one
	if err := initializer.Db.Where("user_id = ?", user.ID).First(&userOtp).Error; err != nil {
		if err.Error() == "record not found" {
			userOtp.Attempt = 1
			userOtp.UserId = user.ID
			userOtp.OtpCode = code
			userOtp.IsUsed = false
			if err := initializer.Db.Create(&userOtp).Error; err != nil {
				return C.Status(500).JSON(fiber.Map{
					"message": err.Error(),
				})
			}
			return C.Status(201).JSON(fiber.Map{
				"message": "Otp sent",
				"Otp":     code,
			})
		}
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userOtp.Attempt += 1
	userOtp.OtpCode = code
	userOtp.IsUsed = false

	if err := initializer.Db.Save(&userOtp).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(200).JSON(fiber.Map{
		"message": "Otp sent",
		"Otp":     code,
	})
}

// verify otp for authentication
type VeryfyOtpBody struct {
	Email string
	Otp   int
}

func VeryfyOtp(C *fiber.Ctx) error {
	var body VeryfyOtpBody
	var user model.User

	if err := C.BodyParser(&body); err != nil {
		return C.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := initializer.Db.Where("email = ?", body.Email).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return C.Status(404).JSON(fiber.Map{
				"message": "Email:" + err.Error(),
			})
		}
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var userOtp model.User_Otp
	if err := initializer.Db.Where("user_id = ? and otp_code = ?", user.ID, body.Otp).First(&userOtp).Error; err != nil {
		if err.Error() == "record not found" {
			return C.Status(404).JSON(fiber.Map{
				"message": "Invalid Otp",
			})
		}
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if userOtp.IsUsed {
		return C.Status(400).JSON(fiber.Map{
			"message": "Otp Veryfied already ",
		})
	}

	if !userOtp.IsUsed {
		userOtp.IsUsed = true
		userOtp.Attempt = 0
		user.IsEmailVerified = true
		if err := initializer.Db.Save(&userOtp).Error; err != nil {
			return C.Status(500).JSON(fiber.Map{
				"message": "Otp: " + err.Error(),
			})
		}
		if err := initializer.Db.Save(&user).Error; err != nil {
			return C.Status(500).JSON(fiber.Map{
				"message": "User: " + err.Error(),
			})
		}
	}

	userType := "member"
	if user.IsListner {
		userType = "listner"
	}
	token, err := jwtToken.GenerateToken(user.ID, userType)
	if err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": "token: " + err.Error(),
		})
	}

	if userType == "listner" {
		userType = "listener"
	}

	return C.Status(200).JSON(fiber.Map{
		"message": "User Veryfied successfully",
		"token":   token,
		"data":    userType,
	})
}
