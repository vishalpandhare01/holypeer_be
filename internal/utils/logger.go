package utils

import (
	"math/rand"
	"time"
)

func Otp_Number_Generate() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(900000) + 100000
	return randomNumber
}
