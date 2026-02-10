package http

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-hexagonal-practice/internal/core/ports"
// )

// type AuthHandler struct {
// 	authService ports.AuthServicePort
// }

// func NewAuthHandler(as ports.AuthServicePort) *AuthHandler {
// 	return &AuthHandler{authService: as}
// }

// // LoginRequest is the DTO for incoming JSON
// type LoginRequest struct {
// 	Email    string `json:"email" binding:"required,email"`
// 	Password string `json:"password" binding:"required"`
// }

// func (h *AuthHandler) Login(c *gin.Context) {
// 	var req LoginRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
// 		return
// 	}

// 	// Extract Metadata from request headers
// 	meta := ports.SessionMetadata{
// 		IPAddress: c.ClientIP(),
// 		UserAgent: c.Request.UserAgent(),
// 	}

// 	// Call the Core Service
// 	session, token, err := h.authService.Login(c.Request.Context(), req.Email, req.Password, meta)
// 	if err != nil {
// 		// Log the actual error internally and return a generic message
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Optionally set the token as a cookie or return in JSON
// 	c.SetCookie("refresh_token", token, 3600*24, "/", "", true, true)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message":    "login successful",
// 		"session_id": session.ID,
// 		"token":      token, // Usually you'd send an Access Token here too
// 	})
// }
