package jwtToken

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenMap struct {
	UserId   string
	UserTypr string
	Token    string
}

func GenerateToken(UserId string, UserTypr string) (string, error) {
	secretKey := os.Getenv("SECREAT_KEY")
	secretKeyBytes := []byte(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    UserId,
		"user_type": UserTypr,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(secretKeyBytes)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
