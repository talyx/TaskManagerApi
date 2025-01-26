package middleware

import (
	"github.com/talyx/TaskManagerApi/internal/services"
	"github.com/talyx/TaskManagerApi/pkg/logger"
	"net/http"
)

func AuthMiddleware(authService *services.AuthService) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := authService.Athenticate(w, r)
			if err != nil {
				logger.Error("Unauthorized ", map[string]interface{}{"error": err.Error()})
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			logger.Info("Authenticated user", nil)
			handler.ServeHTTP(w, r)
		})
	}

}
