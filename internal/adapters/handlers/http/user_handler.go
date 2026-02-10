package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-hexagonal-practice/internal/core/ports"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type userHandler struct {
	userService ports.UserUseCase
}

func NewUserHandler(service ports.UserUseCase) *userHandler {
	return &userHandler{
		userService: service,
	}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var req UserAccountRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeJSONError(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}

	// Validate Fields
	if err := validate.Struct(req); err != nil {
		h.writeJSONError(w, http.StatusBadRequest, "Missing or invalid required fields "+err.Error())
		return
	}

	// if (req.CountryCode != nil && req.CountrySource == nil) || (req.CountrySource != nil && req.CountryCode == nil) {
	// 	h.writeJSONError(w, http.StatusBadRequest, "Country code and source both should be provded or vise versa")
	// 	return
	// }
	defer r.Body.Close()

	// ipAddress := strings.Split(r.RemoteAddr, ":")[0]
	ipAddress := r.RemoteAddr
	if strings.Contains(ipAddress, "[::1]") {
		ipAddress = "127.0.0.1"
	} else {
		ipAddress = strings.Split(ipAddress, ":")[0]
	}
	userAgent := h.stringPtr(r.Header.Get("User-Agent"))
	// sessionInit := domain_sessions.UserSessions{
	// 	IPAddress: ipAddress,
	// 	UserAgent: h.stringPtr(r.Header.Get("User-Agent")),
	// }
	// userProfileInit := domain_profile.UserProfiles{
	// 	UserID: ,
	// }

	ctx := r.Context()
	sessionRecord, err := h.userService.Register(ctx, req.Email, req.Password, ipAddress, userAgent)

	if err != nil {
		h.writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response := UserAccountRegisterResponse{
		SessionID:    sessionRecord.ID.String(),
		UserID:       sessionRecord.UserID.String(),
		RefreshToken: sessionRecord.RefreshTokenHash,
		ExpiresAt:    sessionRecord.ExpiresAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *userHandler) writeJSONError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *userHandler) stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
