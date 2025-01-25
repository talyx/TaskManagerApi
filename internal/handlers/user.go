package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/talyx/TaskManagerApi/internal/models"
	"github.com/talyx/TaskManagerApi/internal/services"
	"github.com/talyx/TaskManagerApi/pkg/logger"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	logger.Debug("body:", map[string]interface{}{
		"body": r.Body,
	})
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Error("Invalid request body", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	createUser, err := h.UserService.CreateUser(&user)
	if err != nil {
		logger.Error("Error creating user", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createUser)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error("Invalid id parameter", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	user, err := h.UserService.GetUserById(uint(id))
	if err != nil {
		logger.Error("Error getting user", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}
	if user == nil {
		logger.Error("user not found", nil)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error("Invalid id parameter", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Error("Invalid request body", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	updateUser, err := h.UserService.UpdateUser(uint(id), user.Names, user.Email)
	if err != nil {
		logger.Error("Error updating user", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updateUser)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logger.Error("Invalid id parameter", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}
	err = h.UserService.DeleteUserById(uint((id)))
	if err != nil {
		logger.Error("Error deleting user", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		logger.Error("Error getting all user", map[string]interface{}{
			"error": err,
		})
		http.Error(w, "Error getting all user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)

}
