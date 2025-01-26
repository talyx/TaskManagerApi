package handlers

import (
	"encoding/json"
	"github.com/talyx/TaskManagerApi/internal/services"
	"github.com/talyx/TaskManagerApi/pkg/logger"
	"net/http"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Invalid login request", map[string]interface{}{"error": err.Error()})
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	err := handler.AuthService.Login(w, r)
	if err != nil {
		logger.Error("Authentication error", map[string]interface{}{"error": err.Error()})
		http.Error(w, "Authentication error", http.StatusUnauthorized)
		return
	}
	logger.Info("Authentication success", map[string]interface{}{"user": req.Login})
}

func (handler *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	err := handler.AuthService.Logout(w, r)
	if err != nil {
		logger.Error("Logout error", map[string]interface{}{"error": err.Error()})
		http.Error(w, "Logout error", http.StatusUnauthorized)
		return
	}
	logger.Info("Logout success", map[string]interface{}{"user": nil})
	w.WriteHeader(http.StatusOK)
}
