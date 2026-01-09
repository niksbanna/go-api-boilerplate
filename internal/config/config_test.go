package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Test with default values
	config, err := Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if config.Server.Port != "3000" {
		t.Errorf("Expected default port 3000, got %s", config.Server.Port)
	}

	if config.Server.Host != "0.0.0.0" {
		t.Errorf("Expected default host 0.0.0.0, got %s", config.Server.Host)
	}
}

func TestLoadWithEnv(t *testing.T) {
	// Set environment variables
	if err := os.Setenv("SERVER_PORT", "8080"); err != nil {
		t.Fatalf("Failed to set SERVER_PORT: %v", err)
	}
	if err := os.Setenv("SERVER_HOST", "localhost"); err != nil {
		t.Fatalf("Failed to set SERVER_HOST: %v", err)
	}
	defer func() {
		_ = os.Unsetenv("SERVER_PORT") //nolint:errcheck
		_ = os.Unsetenv("SERVER_HOST") //nolint:errcheck
	}()

	config, err := Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if config.Server.Port != "8080" {
		t.Errorf("Expected port 8080, got %s", config.Server.Port)
	}

	if config.Server.Host != "localhost" {
		t.Errorf("Expected host localhost, got %s", config.Server.Host)
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "Valid config",
			config: &Config{
				Server: ServerConfig{
					Port: "3000",
					Host: "0.0.0.0",
					Env:  "development",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid port - empty",
			config: &Config{
				Server: ServerConfig{
					Port: "",
					Host: "0.0.0.0",
					Env:  "development",
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid port - not a number",
			config: &Config{
				Server: ServerConfig{
					Port: "abc",
					Host: "0.0.0.0",
					Env:  "development",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetDatabaseDSN(t *testing.T) {
	config := &Config{
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "secret",
			DBName:   "testdb",
			SSLMode:  "disable",
		},
	}

	expected := "host=localhost port=5432 user=postgres password=secret dbname=testdb sslmode=disable"
	dsn := config.GetDatabaseDSN()

	if dsn != expected {
		t.Errorf("GetDatabaseDSN() = %v, want %v", dsn, expected)
	}
}
