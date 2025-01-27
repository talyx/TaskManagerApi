package handlers

import (
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
	err := handler.AuthService.Login(w, r)
	if err != nil {
		logger.Error("Authorization error", map[string]interface{}{"error": err.Error()})
		http.Error(w, "Authorization error", http.StatusUnauthorized)
		return
	}
	logger.Info("Authentication success", nil)
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
