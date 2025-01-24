package config_test

import (
	"github.com/talyx/TaskManagerApi/internal/config"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("SERVER_PORT", "9000")
	defer os.Unsetenv("SERVER_PORT")

	cfg := config.LoadConfig()
	if cfg.ServerPort != "9000" {
		t.Errorf("expected SERVER_PORT to be 9000, got %s", cfg.ServerPort)
	}
	if cfg.DBHost != "localhost" {
		t.Errorf("expected DBHost to be localhost, got %s", cfg.DBHost)
	}
}
