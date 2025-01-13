package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string `gorm:"primaryKey;type:char(36);not null"`
	Name        string `gorm:"type:varchar(255);not:null"`
	Email       string `gorm:"type:varchar(255);not:null"`
	DateOfBirth string `gorm:"type:varchar(255);not:null"`
	// Password        string    `gorm:"type:varchar(255);not:null"`
	IsBlock         bool      `gorm:"type:boolean;default:false"`
	IsEmailVerified bool      `gorm:"type:boolean;default:false"`
	IsListner       bool      `gorm:"default:false"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

type User_Otp struct {
	ID        string    `gorm:"primaryKey;type:char(36);not null"`
	OtpCode   int       `gorm:"type:char(6);not:null"`
	IsUsed    bool      `gorm:"bool;default:false"`
	UserId    string    `gorm:"type:varchar(36);not:null"`
	Attempt   int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (b *User) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}

func (b *User_Otp) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}
