package services

import (
	"github.com/talyx/TaskManagerApi/internal/auth"
	"net/http"
)

type AuthService struct {
	Authenticator auth.Authenticator
}

func NewAuthService(authenticator auth.Authenticator) *AuthService {
	return &AuthService{authenticator}
}

func (as *AuthService) Login(w http.ResponseWriter, r *http.Request) error {
	return as.Authenticator.Authorize(w, r)
}

func (as *AuthService) Logout(w http.ResponseWriter, r *http.Request) error {
	return as.Authenticator.Logout(w, r)
}

func (as *AuthService) Athenticate(w http.ResponseWriter, r *http.Request) error {
	return as.Authenticator.Authenticate(w, r)

}

func (as *AuthService) GetUserID(w http.ResponseWriter, r *http.Request) (uint, error) {
	return as.Authenticator.GetUserID(w, r)
}
