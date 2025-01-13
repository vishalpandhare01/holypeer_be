package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemeberSchema struct {
	ID            string    `gorm:"type:char(36);primarykey;not null"`
	UserID        string    `gorm:"type:char(36);not null"`
	TodysFeel     string    `gorm:"type:enum('depression','anxiety','relationships','ocd','parenting','family','loneliness','happy','good');not null"`
	ChatKey       string    `gorm:"type:text"`
	IsChatKeyUsed bool      `gorm:"type:boolean"`
	Bio           string    `gorm:"type:text"`
	Country       string    `gorm:"type:varchar(255); not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
	User          *User     `gorm:"foreignKey:UserID;constraint;onDelete:CASCADE"`
}

func (M *MemeberSchema) BeforeCreate(tx *gorm.DB) (err error) {
	if M.ID == "" {
		M.ID = uuid.New().String()
	}
	return
}
