package auth

import "net/http"

type Authenticator interface {
	Authorize(w http.ResponseWriter, r *http.Request) error
	Logout(w http.ResponseWriter, r *http.Request) error
	Authenticate(w http.ResponseWriter, r *http.Request) error
	GetUserID(w http.ResponseWriter, r *http.Request) (uint, error)
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
