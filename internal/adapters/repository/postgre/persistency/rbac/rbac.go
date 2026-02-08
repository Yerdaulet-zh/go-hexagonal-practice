package rbac

import (
	"time"

	user_model "github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/google/uuid"
)

/*
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL, -- 'admin', 'user', 'auditor'
    description TEXT
);

CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    slug VARCHAR(50) UNIQUE NOT NULL, -- 'users.read', 'users.write', 'audit.view'
    description TEXT
);

-- Join table for Role-Permissions
CREATE TABLE role_permissions (
    role_id INT REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INT REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

-- Assign roles to users
CREATE TABLE user_roles (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role_id INT REFERENCES roles(id) ON DELETE CASCADE,
    assigned_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    assigned_by UUID REFERENCES users(id), -- Who gave this role? Audit trail.
    PRIMARY KEY (user_id, role_id)
);
*/

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

	User *user_model.User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Role *Role            `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
