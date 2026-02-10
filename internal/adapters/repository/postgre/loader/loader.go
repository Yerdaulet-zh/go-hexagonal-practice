package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/oauth"
	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/profile"
	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/rbac"
	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/sessions"
	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/user"
	"github.com/go-hexagonal-practice/internal/adapters/repository/postgre/persistency/verification"

	// "github.com/go-hexagonal-practice/internal/core/domain/user_test/user"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // This is the magic switch
		},
	}
	// Atlas needs to see the physical Go structs to generate the SQL
	stmts, err := gormschema.New("postgres", gormschema.WithConfig(config)).Load(
		&user.User{},
		&user.UserCredentials{},
		&oauth.OauthIdentities{},
		&profile.UserProfile{},
		&profile.UserProfileHistory{},
		&sessions.UserSessions{},
		&verification.Verifications{},
		&verification.UsersMFASecrets{},
		&rbac.Role{},
		&rbac.RolePermissions{},
		&rbac.Permissions{},
		&rbac.UserRoles{},
		// &user.Session{},
		// &user.UserNameHistory{},
		// &user.UserSurnameHistory{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
