package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserData struct {
	UserId   string
	UserType string
}

func verifyToken(tokenString string) (*UserData, error) {
	secretKey := os.Getenv("SECREAT_KEY")
	secretKeyBytes := []byte(secretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKeyBytes, nil
	})

	if err != nil {
		return nil, err
	}

	var UserData UserData
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Now you can access the claim like a map
		userId, userok := claims["userId"].(string)
		userTpe, userTypeok := claims["user_type"].(string)
		if userok && userTypeok {
			UserData.UserId = userId
			UserData.UserType = userTpe
			// fmt.Println("UserID:", userId, "usertype: ", userTpe)
		} else {
			fmt.Println("UserID claim not found or invalid type")
			return nil, fmt.Errorf("invalid token")
		}
	} else {
		fmt.Println("Invalid token")
	}
	return &UserData, nil
}

func Authentication(C *fiber.Ctx) error {
	tokenString := C.Get("Authorization")

	if tokenString == "" {
		return C.Status(403).JSON(fiber.Map{
			"message": "You are not authorized , authorize header missing",
		})
	}

	tokenString = tokenString[len("Bearer "):]

	userdata, err := verifyToken(tokenString)
	if err != nil {
		return C.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	C.Locals("userId", userdata.UserId)
	C.Locals("userType", userdata.UserType)
	return C.Next()
}
