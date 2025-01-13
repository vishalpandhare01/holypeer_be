package migration

import (
	"fmt"
	"log"

	"github.com/vishalpandhare01/holypeer_backend/initializer"
	"github.com/vishalpandhare01/holypeer_backend/internal/model"
)

func DbMigration() {
	err := initializer.Db.AutoMigrate(
		&model.User{},
		&model.User_Otp{},
		&model.MemeberSchema{},
		&model.ListenerSchema{},
	)

	if err != nil {
		log.Fatal("Migration failed 🤣: ", err)
	}
	fmt.Println("All table Migrate successfully!! 🚀 ")
}
