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
	err := as.Authenticator.Authorize(w, r)
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthService) Logout(w http.ResponseWriter, r *http.Request) error {
	err := as.Authenticator.Logout(w, r)
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthService) Athenticate(w http.ResponseWriter, r *http.Request) error {
	err := as.Authenticator.Authenticate(w, r)
	if err != nil {
		return err
	}
	return nil
}
