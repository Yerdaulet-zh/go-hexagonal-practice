package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/badoux/checkmail"
	domain_profile "github.com/go-hexagonal-practice/internal/core/domain/profile"
	domain_sessions "github.com/go-hexagonal-practice/internal/core/domain/sessions"
	domain_user "github.com/go-hexagonal-practice/internal/core/domain/user"
	"github.com/go-hexagonal-practice/internal/core/ports"
	"golang.org/x/crypto/argon2"
)

type UserService struct {
	userRepo ports.UserPort
}

func NewUserService(repo ports.UserPort) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (s *UserService) Register(
	ctx context.Context,
	params ports.RegisterParams,
) (*domain_sessions.UserSessions, error) {
	if err := checkmail.ValidateFormat(params.Email); err != nil {
		return nil, fmt.Errorf("Email syntax validation error: %s", err.Error())
	}
	if err := checkmail.ValidateHost(params.Email); err != nil {
		return nil, fmt.Errorf("Email service validation error: %s", err.Error())
	}

	existingUser, err := s.userRepo.GetUserByEmail(ctx, params.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user already exists")
	}
	// if err != nil {
	// 	return nil, fmt.Errorf("Internal server error: %w", err)
	// }

	hashedPassword, passwordSalt, err := hashPassword(params.Password, "")
	if err != nil {
		return nil, err
	}
	newUser := &domain_user.User{
		Email:      params.Email,
		UserStatus: "pending_verification",
	}

	userCredentials := &domain_user.UserCredentials{
		PasswordHash: hashedPassword,
		PasswordSalt: passwordSalt,
	}

	// UserSession
	token, err := generateSecureToken()
	if err != nil {
		return nil, err
	}
	expiration := time.Now().Add(24 * time.Hour)
	userSessionRecord := domain_sessions.UserSessions{
		RefreshTokenHash: hashToken(token),
		IPAddress:        params.IPAddress,
		UserAgent:        params.UserAgent,
		Device:           params.Device,
		ExpiresAt:        expiration,
	}

	userProfileRecord := domain_profile.UserProfile{
		FirstName:     params.FirstName,
		LastName:      params.LastName,
		CountryCode:   params.CountryCode,
		CountrySource: params.CountrySource,
	}

	session, err := s.userRepo.CreateUser(ctx, newUser, userCredentials, &userSessionRecord, &userProfileRecord)
	if err != nil {
		return nil, err
	}
	session.RefreshTokenHash = token

	return session, nil
}

func hashPassword(password string, existingSalt string) (string, string, error) {
	var saltBytes []byte
	var err error

	if existingSalt != "" {
		// --- Login Flow ---
		saltBytes, err = base64.RawStdEncoding.DecodeString(existingSalt)
		if err != nil {
			return "", "", fmt.Errorf("Error while decoding existing salt %w", err)
		}
	} else {
		// --- Registration flow ---
		// Generate a new random 16-byte salt
		saltBytes = make([]byte, 16)
		if _, err = rand.Read(saltBytes); err != nil {
			return "", "", fmt.Errorf("failed to generate salt: %w", err)
		}
	}

	hashBytes := argon2.IDKey([]byte(password), saltBytes, 1, 64*1024, 4, 32)

	saltString := base64.RawStdEncoding.EncodeToString(saltBytes)
	hashString := base64.RawStdEncoding.EncodeToString(hashBytes)

	return hashString, saltString, nil
}

func generateSecureToken() (string, error) {
	// 32 bytes of randomness is plenty (256 bits of entropy)
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// Use RawURLEncoding to avoid '+' or '/' which can be messy in headers/URLs
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
