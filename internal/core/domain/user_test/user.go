package user_test

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email    string    `gorm:"type:varchar(255);unique"`
	Password string    `gorm:"type:text;not null"`
	Name     *string   `gorm:"type:varchar(100)"`
	Surname  *string   `gorm:"type:varchar(100)"`
	Country  *string   `gorm:"type:varchar(100)"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// ORM-only relations (NO SCHEMA)
	NameHistory    []UserNameHistory    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	SurnameHistory []UserSurnameHistory `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type UserNameHistory struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;index:idx_user_temporal;not null"`

	Name string `gorm:"type:varchar(100);not null"`
	// Surname string `gorm:"type:varchar(100);not null"`

	// Temporal tracking (SCD Type 4)
	ValidFrom time.Time  `gorm:"index:idx_user_temporal;not null"`
	ValidTo   *time.Time `gorm:"index:idx_user_temporal"` // NULL means "Current"

	IPv4Address string         `gorm:"type:varchar(45)"`
	UserAgent   string         `gorm:"type:text"`
	ChangedBy   uuid.UUID      `gorm:"type:uuid;not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Many-to-one relations
	User User `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type UserSurnameHistory struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;index:idx_user_temporal;not null"`

	// Name    string `gorm:"type:varchar(100);not null"`
	Surname string `gorm:"type:varchar(100);not null"`

	// Temporal tracking (SCD Type 4)
	ValidFrom time.Time  `gorm:"index:idx_user_temporal;not null"`
	ValidTo   *time.Time `gorm:"index:idx_user_temporal"` // NULL means "Current"

	IPv4Address string         `gorm:"type:varchar(45)"`
	UserAgent   string         `gorm:"type:text"`
	ChangedBy   uuid.UUID      `gorm:"type:uuid;not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Many-to-one relations
	User User `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
