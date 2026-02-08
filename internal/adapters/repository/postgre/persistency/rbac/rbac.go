package rbac

import (
	"time"

	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/user"
	"github.com/google/uuid"
)

type Role struct {
	ID          uint64  `gorm:"primaryKey;autoIncrement:true"`
	Name        string  `gorm:"type:varchar(50);unique;not null"` // admin, user, auditor
	Description *string `gorm:"type:text"`
}

type Permissions struct {
	ID          uint64  `gorm:"primaryKey;autoIncrement:true"`
	Slug        string  `gorm:"type:varchar(50);unique;not null"` // users.read, users.write, audit.view
	Description *string `gorm:"type:text"`
}

type RolePermissions struct {
	RoleID       uint64 `gorm:"primaryKey"`
	PermissionID uint64 `gorm:"primaryKey"`

	Role        *Role        `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Permissions *Permissions `gorm:"foreignKey:PermissionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type UserRoles struct {
	UserID     uuid.UUID  `gorm:"type:uuid;primaryKey"`
	RoleID     uint64     `gorm:"primaryKey"`
	AssignedAt time.Time  `gorm:"type:timestamptz;default:now();not null"`
	AssignedBy *uuid.UUID `gorm:"type:uuid;not null"` // NULL if assigned by a system

	User *user.User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Role *Role      `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
