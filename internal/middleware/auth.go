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
	// First, try to get the token from the Authorization header
	tokenString := C.Get("Authorization")
	if tokenString == "" {
		// If no token in the header, check for the token in the query string
		tokenString = C.Query("Authorization")
	}

	// If no token is found at all
	if tokenString == "" {
		return C.Status(403).JSON(fiber.Map{
			"message": "You are not authorized, Authorization header or query parameter missing",
		})
	}

	// Strip the "Bearer " prefix if it's present
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Verify the token
	userdata, err := verifyToken(tokenString)
	if err != nil {
		return C.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Set user data in the context for downstream handlers
	C.Locals("userId", userdata.UserId)
	C.Locals("userType", userdata.UserType)
	return C.Next()
}
