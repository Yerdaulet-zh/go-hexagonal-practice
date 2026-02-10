package http

import "time"

type UserAccountRegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required"`

	LastName      *string `json:"last_name,omitempty"`
	CountryCode   *string `json:"country_code,omitempty" validate:"required_with=CountrySource"`
	CountrySource *string `json:"country_source,omitempty" validate:"required_with=CountryCode"`
}

type UserAccountRegisterResponse struct {
	SessionID    string    `json:"session_id"`
	UserID       string    `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}
