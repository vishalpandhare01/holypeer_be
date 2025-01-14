package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	ID         string          `gorm:"type:char(36);primarykey;not null"`
	ListenerID *string         `gorm:"type:char(36);"`
	MemberID   string          `gorm:"type:char(36);not null"`
	IsAccepted bool            `gorm:"default:false"`
	ChatKey    string          `gorm:"type:text"`
	Member     *MemeberSchema  `gorm:"foreignKey:MemberID;"`
	Listener   *ListenerSchema `gorm:"foreignKey:ListenerID;"`
	CreatedAt  time.Time       `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt  time.Time       `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (F *Chat) BeforeCreate(tx *gorm.DB) (err error) {
	if F.ID == "" {
		F.ID = uuid.New().String()
	}
	return
}
