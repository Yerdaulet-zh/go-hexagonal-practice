package rbac

import (
	"time"

	user_model "github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/google/uuid"
)

type Role struct {
	ID          uint64
	Name        string // admin, user, auditor
	Description *string
}

type Permissions struct {
	ID          uint64
	Slug        string // users.read, users.write, audit.view
	Description *string
}

type RolePermissions struct {
	RoleID       uint64
	PermissionID uint64

	Role        *Role
	Permissions *Permissions
}

type UserRoles struct {
	UserID     uuid.UUID
	RoleID     uint64
	AssignedAt time.Time
	AssignedBy *uuid.UUID // NULL if assigned by a system

	User *user_model.User
	Role *Role
}
