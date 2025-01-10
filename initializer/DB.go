package initializer

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() {
	var error error
	var dbUrl = os.Getenv("DATABASE_URL")
	Db, error = gorm.Open(mysql.Open(dbUrl))
	if error != nil {
		log.Fatal("Database connnectiona failed 🤣: ", error)
	}

}
