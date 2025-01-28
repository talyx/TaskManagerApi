package auth

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/talyx/TaskManagerApi/internal/database"
	"github.com/talyx/TaskManagerApi/internal/utils"
	"github.com/talyx/TaskManagerApi/pkg/logger"
	"net/http"
	"strings"
	"time"
)

type JWTAuth struct {
	secretString string
	UserRepo     *database.UserRepository
}

//const sectretString = "secretCode"

func (j *JWTAuth) Authorize(w http.ResponseWriter, r *http.Request) error {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	user, err := j.UserRepo.GetUserByLogin(req.Login)
	if err != nil {
		return errors.New("user not found")
	}
	if err = utils.ComparePassword(user.PasswordHash, req.Password); err != nil {
		return errors.New("invalid password")
	}
	token, err := j.GenerateJWT(user.ID)
	if err != nil {
		return errors.New("failed to generate token")
	}
	response := map[string]string{"token": token}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return nil
}

func (j *JWTAuth) Logout(w http.ResponseWriter, r *http.Request) error {
	logger.Error("logout jwt not relized", map[string]interface{}{
		"response": w,
		"request":  r,
	})
	return errors.New("jwt logout not implemented")
}

func (j *JWTAuth) GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretString))
}

func (j *JWTAuth) Authenticate(w http.ResponseWriter, r *http.Request) error {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return errors.New("no auth header")
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Error("unexpected signing method", map[string]interface{}{
				"token": tokenString,
				"error": ok,
			})
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretString), nil
	})
	if err != nil {
		logger.Error("failed to parse token", map[string]interface{}{
			"token":       tokenString,
			"error":       err,
			"tokenString": tokenString,
		})
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return errors.New("invalid token")
	}
	_, ok = claims["user_id"].(float64)
	if !ok {
		return errors.New("invalid token")
	}
	return nil
}

func (j *JWTAuth) GetUserID(w http.ResponseWriter, r *http.Request) (uint, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {

		tokenString := strings.TrimPrefix(authHeader, "Bearer")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.Error("unexpected signing method", map[string]interface{}{
					"token": tokenString,
					"error": ok,
				})
				return nil, errors.New("unexpected signing method")
			}
			return []byte(j.secretString), nil
		})
		if err != nil {
			logger.Error("failed to parse token", map[string]interface{}{
				"token": tokenString,
				"error": err,
			})
			return 0, err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return 0, errors.New("invalid token")
		}
		userID, ok := claims["userID"].(float64)
		if !ok {
			return 0, errors.New("invalid token")
		}
		return uint(userID), nil
	}
	return 0, errors.New("invalid token")
}

func NewJwtAuth(s string, repo *database.UserRepository) *JWTAuth {
	return &JWTAuth{secretString: s, UserRepo: repo}
}
