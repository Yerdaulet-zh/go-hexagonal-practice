package user_test

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;index:idx_session_user;not null"`

	IPv4Address string `gorm:"type:varchar(45)"`
	UserAgent   string `gorm:"type:text"`

	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"column:last_active_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relations
	User User `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
