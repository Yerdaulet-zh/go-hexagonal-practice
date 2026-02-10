package profile

import (
	"time"

	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/user"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type UserProfile struct {
	UserID        uuid.UUID       `gorm:"type:uuid;primaryKey"`
	FirstName     string          `gorm:"type:varchar(100);not null"`
	LastName      *string         `gorm:"type:varchar(100)"`
	CountryCode   *string         `gorm:"type:varchar(2)"`
	CountrySource *string         `gorm:"type:varchar(50)"`
	AvatarURL     *string         `gorm:"type:text"`
	Preferences   *datatypes.JSON `gorm:"type:json"`
	UpdatedAt     time.Time       `gorm:"type:timestamptz;default:now();not null"`

	User *user.User `gorm:"foreignKey:UserID;reference:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type UserProfileHistory struct {
	HistoryId     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID        uuid.UUID `gorm:"type:uuid;not null"`
	FirstName     string    `gorm:"type:varchar(100);not null"`
	LastName      *string   `gorm:"type:varchar(100)"`
	CountryCode   *string   `gorm:"type:varchar(2)"`
	CountrySource *string   `gorm:"type:varchar(50)"`
	AvatarURL     *string   `gorm:"type:text"`
	ChangedAt     time.Time `gorm:"type:timestamptz;default:now();not null"`
	Operation     string    `gorm:"type:varchar(10);not null"`

	User *user.User `gorm:"foreignKey:UserID;reference:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}
