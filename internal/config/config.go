package config

import (
	"github.com/joho/godotenv"
	"github.com/talyx/TaskManagerApi/pkg/logger"
	"os"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	LogLevel   string
	LogOutput  string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		logger.Error("No .env file found, falling back to env variables", nil)
	}

	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "user"),
		DBPassword: getEnv("DB_PASSWORD", "admin"),
		DBName:     getEnv("DB_NAME", "mydb"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
		LogOutput:  getEnv("LOG_OUTPUT", ""),
	}
}

func (cfg *Config) PrintConfig() {
	logger.Debug("Config data:", map[string]interface{}{
		"server_port": cfg.ServerPort,
		"db_host":     cfg.DBHost,
		"db_port":     cfg.DBPort,
		"db_user":     cfg.DBUser,
		"db_password": cfg.DBPassword,
		"db_name":     cfg.DBName,
		"log_level":   cfg.LogLevel,
		"log_output":  cfg.LogOutput,
	})
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
