package app

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/talyx/TaskManagerApi/internal/auth"
	"github.com/talyx/TaskManagerApi/internal/config"
	"github.com/talyx/TaskManagerApi/internal/database"
	"github.com/talyx/TaskManagerApi/internal/handlers"
	"github.com/talyx/TaskManagerApi/internal/middleware"
	"github.com/talyx/TaskManagerApi/internal/services"
	"github.com/talyx/TaskManagerApi/pkg/logger"
	"log"
	"net/http"
)

func Run() error {
	// init logger
	if err := initLogger(); err != nil {
		return err
	}
	cfg := config.LoadConfig() //Load config
	database.InitDatabase(cfg) // Connection to db

	router, err := initRouts()
	if err != nil {
		return err
	}
	logger.Info("Server started", map[string]interface{}{
		"version":    "1.0.0",
		"serverPort": cfg.ServerPort,
	})
	return http.ListenAndServe(fmt.Sprintf(":%s", cfg.ServerPort), router) // run server
}

func initLogger() error {
	logLevel := flag.String("log-level", "info", "Уровень логирования: debug, info, warn, error")
	logOutput := flag.String("log-output", "", "Путь к файлу лога (по умолчанию вывод в консоль)")
	flag.Parse()
	err := logger.InitLogger(*logLevel, *logOutput)
	if err != nil {
		log.Fatalf("Log init fatal error: %v", err)
	}
	return nil
}

func initRouts() (*mux.Router, error) {
	// 		Repositories
	userRepo := database.NewUserRepository(database.DB)
	projectRepo := database.NewProjectRepository(database.DB)
	taskRepo := database.NewTaskRepository(database.DB)

	sessionStore := sessions.NewFilesystemStore("cookies", []byte("your-secret-key"))

	//		Services
	userService := services.NewUserService(userRepo)
	projectService := services.NewProjectService(projectRepo)
	taskService := services.NewTaskService(taskRepo, projectRepo)
	//jwtAuth := auth.NewJwtAuth("secretCode", userRepo)
	sessionAuth := auth.NewSessionAuth(sessionStore, userRepo)
	authService := services.NewAuthService(sessionAuth)

	//		Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	projectHandler := handlers.NewProjectHandler(projectService, sessionStore)
	taskHandler := handlers.NewTaskHandler(taskService, sessionStore)

	// 		router
	router := mux.NewRouter()
	protected := router.PathPrefix("/").Subrouter()

	//		Authorization routs
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/logout", authHandler.Login).Methods("POST")
	router.HandleFunc("/register", userHandler.CreateUser).Methods("POST")

	// 		middleware
	protected.Use(middleware.AuthMiddleware(authService))

	//		User routs
	protected.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	protected.HandleFunc("/user/{id:[0-9]+}", userHandler.GetUserById).Methods("GET")
	protected.HandleFunc("/users/{id:[0-9]+}", userHandler.UpdateUser).Methods("PUT")
	protected.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUser).Methods("DELETE")
	protected.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")

	//		Project routs
	protected.HandleFunc("/project", projectHandler.CreateProject).Methods("POST")
	protected.HandleFunc("/project/{id:[0-9]+}", projectHandler.GetProjectById).Methods("GET")
	protected.HandleFunc("/project/{id:[0-9]+}", projectHandler.UpdateProject).Methods("PUT")
	protected.HandleFunc("/project/{id:[0-9]+}", projectHandler.DeleteProject).Methods("DELETE")
	protected.HandleFunc("/projects", projectHandler.GetAllProjectByUserId).Methods("GET")

	//		Task routs
	protected.HandleFunc("/task", taskHandler.CreateTask).Methods("POST")
	protected.HandleFunc("/task/{id:[0-9]+}", taskHandler.GetTaskById).Methods("GET")
	protected.HandleFunc("/task/{id:[0-9]+}", taskHandler.UpdateTask).Methods("PUT")
	protected.HandleFunc("/task/{id:[0-9]+}", taskHandler.DeleteTask).Methods("DELETE")
	protected.HandleFunc("/tasks/{id:[0-9]+}", taskHandler.GetAllTasksByProjectId).Methods("GET")
	// Подмаршрут для защищённых маршрутов

	return router, nil
}
