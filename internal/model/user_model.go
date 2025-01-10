package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          string `gorm:"primaryKey;type:char(36);not:null"`
	Name        string `gorm:"type:varchar(255);not:null"`
	Email       string `gorm:"type:varchar(255);not:null"`
	DateOfBirth string `gorm:"type:varchar(255);not:null"`
	// Password        string    `gorm:"type:varchar(255);not:null"`
	IsBlock         bool      `gorm:"type:boolean;default:false"`
	IsEmailVerified bool      `gorm:"type:boolean;default:false"`
	IsListner       bool      `gorm:"type:boolean;default:false"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

type User_Otp struct {
	ID        string    `gorm:"primaryKey;type:varchar(36)"`
	OtpCode   int       `gorm:"type:size:6;not:null"`
	IsUsed    bool      `gorm:"bool;default:false"`
	UserId    string    `gorm:"type:varchar(36);not:null"`
	Attempt   int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (b *User) BeforCreate() {
	b.ID = uuid.NewString()
	// if b.Password != "" {
	// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(b.Password), bcrypt.DefaultCost)
	// 	if err != nil {
	// 		fmt.Println("Error hashing password:", err)
	// 	}
	// 	fmt.Println("Hashed password:", string(hashedPassword))

	// 	b.Password = string(hashedPassword)
	// }
}

func (b *User_Otp) BeforCreate() {
	b.ID = uuid.NewString()
}
