package http

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"

	"github.com/go-hexagonal-practice/internal/core/ports"
	"github.com/go-playground/validator/v10"
	"github.com/mileusna/useragent"
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

	ipAddress := h.getClientIP(r)
	userAgent := h.stringPtr(r.Header.Get("User-Agent"))
	device := h.parseDevice(userAgent)

	params := ports.RegisterParams{
		Email:         req.Email,
		Password:      req.Password,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		CountryCode:   req.CountryCode,
		CountrySource: req.CountryCode,
		IPAddress:     ipAddress,
		UserAgent:     userAgent,
		Device:        device,
	}
	ctx := r.Context()
	sessionRecord, err := h.userService.Register(ctx, params)

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

func (h *userHandler) getClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// The first IP in the list is the original client
		return strings.TrimSpace(strings.Split(xff, ",")[0])
	}

	// Check X-Real-IP
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// For local testing without Nginx
	ipAddress, _, _ := net.SplitHostPort(r.RemoteAddr)
	if strings.Contains(ipAddress, "[::1]") {
		return "127.0.0.1"
	}

	return "0.0.0.0"
}

func (s *userHandler) parseDevice(uaString *string) *string {
	if uaString == nil || *uaString == "" {
		return nil
	}

	ua := useragent.Parse(*uaString)

	// Combine OS and Device info for a descriptive string
	deviceInfo := ua.OS + " " + ua.Device
	if ua.Mobile {
		deviceInfo += " (Mobile)"
	} else if ua.Tablet {
		deviceInfo += " (Tablet)"
	} else {
		deviceInfo += " (Desktop)"
	}

	return &deviceInfo
}
