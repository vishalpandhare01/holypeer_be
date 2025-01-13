package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeedBack struct {
	ID                string    `gorm:"type:char(36);primarykey;not null"`
	FeedBackGivenId   string    `gorm:"type:char(36);not null"`
	FeedBackReciverId string    `gorm:"type:char(36);not null"`
	FeedBackMessage   string    `gorm:"type:text"`
	Rating            int       `gorm:"type:size(5);"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (F *FeedBack) BeforeCreate(tx *gorm.DB) (err error) {
	if F.ID == "" {
		F.ID = uuid.New().String()
	}
	return
}
