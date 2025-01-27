package database

import (
	"fmt"
	"github.com/talyx/TaskManagerApi/internal/config"
	"github.com/talyx/TaskManagerApi/internal/models"
	"github.com/talyx/TaskManagerApi/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s port=%s password=%s dbname=%s user=%s sslmode=disable search_path=public",
		cfg.DBHost, cfg.DBPort, cfg.DBPassword, cfg.DBName, cfg.DBUser)
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Connection to database failed", map[string]interface{}{
			"host":     cfg.DBHost,
			"port":     cfg.DBPort,
			"password": cfg.DBPassword,
			"dbname":   cfg.DBName,
			"user":     cfg.DBUser,
			"error":    err,
		})
	}

	logger.Info("Connected to database successfully", nil)
	err = DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Task{})
	if err != nil {
		logger.Error("Migration error: ", map[string]interface{}{
			"error": err,
		})
	}
	logger.Info("Migration successfully", nil)
}
