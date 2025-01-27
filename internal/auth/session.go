package auth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/sessions"
	"github.com/talyx/TaskManagerApi/internal/database"
	"github.com/talyx/TaskManagerApi/internal/utils"
	"github.com/talyx/TaskManagerApi/pkg/logger"
	"net/http"
	"strings"
)

type SessionAuth struct {
	UserRepo     *database.UserRepository
	SessionStore *sessions.FilesystemStore
}

func (s *SessionAuth) Authorize(w http.ResponseWriter, r *http.Request) error {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("http request error", map[string]interface{}{"error": err})
		return err
	}
	user, err := s.UserRepo.GetUserByLogin(req.Login)
	if err != nil {
		logger.Error("get user by login error", map[string]interface{}{"error": err.Error()})
		return err
	}
	err = utils.ComparePassword(strings.TrimSpace(user.PasswordHash), strings.TrimSpace(req.Password))
	if err != nil {
		logger.Error("compare password error", map[string]interface{}{"error": err.Error(),
			"db_pass": user.PasswordHash,
			"req_pas": req.Password,
			"login":   req.Login})
		return err
	}
	session, _ := s.SessionStore.Get(r, "session")
	session.Values["user"] = req.Login
	session.Values["UserID"] = user.ID
	logger.Info("session", map[string]interface{}{"user": req.Login, "user_id": user.ID})
	err = session.Save(r, w)
	if err != nil {
		logger.Error("session save error", map[string]interface{}{"error": err.Error()})
		return err
	}
	return nil
}

func (s *SessionAuth) Logout(w http.ResponseWriter, r *http.Request) error {
	session, _ := s.SessionStore.Get(r, "session")
	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		logger.Error("session logout error", map[string]interface{}{"error": err.Error()})
		return err
	}
	return nil
}

func (s *SessionAuth) Authenticate(w http.ResponseWriter, r *http.Request) error {
	session, _ := s.SessionStore.Get(r, "session")
	user := session.Values["user"]

	if user == nil {
		logger.Error("user is nil", map[string]interface{}{"error": "user is nil"})
		return errors.New("user is unauthorized")
	}
	r = r.WithContext(context.WithValue(r.Context(), "user", user))
	return nil
}

func NewSessionAuth(store *sessions.FilesystemStore, repo *database.UserRepository) *SessionAuth {
	return &SessionAuth{
		SessionStore: store,
		UserRepo:     repo,
	}
}
