package validation

import (
	"fmt"

	"github.com/vishalpandhare01/holypeer_backend/initializer"
	"github.com/vishalpandhare01/holypeer_backend/internal/model"
)

var user model.User

func CheckNameExist(name string) bool {
	if err := initializer.Db.Where("name = ?", name).First(&user).Error; err != nil {
		fmt.Println("Name Not Exist: ", err.Error())
		return false
	}
	return true
}

func CheckEmailExist(email string) bool {
	if err := initializer.Db.Where("email = ?", email).First(&user).Error; err != nil {
		fmt.Println("Email Not Exist: ", err.Error())
		return false
	}
	return true
}
