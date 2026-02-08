package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID
	Email        string
	UserStatus   string
	IsMFAEnabled bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type UserCredentials struct {
	UserID             uuid.UUID
	PasswordHash       string
	PasswordSalt       *string
	LastPasswordChange time.Time
	UpdatedAt          time.Time

	User *User
}
