package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ListenerSchema struct {
	ID         string    `gorm:"type:char(36);primarykey;not null"`
	UserID     string    `gorm:"type:char(36);not null"`
	Bio        string    `gorm:"type:text"`
	Country    string    `gorm:"type:varchar(255); not null"`
	IsVerified bool      `gorm:"type:boolean;default:false"`
	Score      float64   `gorm:"type:decimal(5,2);default:0.00"`
	User       *User     `gorm:"foreignKey:UserID;constraint;onDelete:CASCADE"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (L *ListenerSchema) BeforeCreate(tx *gorm.DB) (err error) {
	if L.ID == "" {
		L.ID = uuid.New().String()
	}
	return
}
