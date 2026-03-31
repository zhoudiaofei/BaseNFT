package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config holds the application configuration. minor comment refresh
type Config struct {
	HTTP struct {
		Port int `yaml:"port" env:"HTTP_PORT" default:"8080"`
	} `yaml:"http"`

	Database struct {
		URL    string `yaml:"url" env:"DATABASE_URL" default:"sqlite://./data/app.db?_foreign_keys=1"`
		Driver string `yaml:"driver" env:"DATABASE_DRIVER" default:"sqlite"`
	} `yaml:"database"`

	Log struct {
		Level string `yaml:"level" env:"LOG_LEVEL" default:"info"`
	} `yaml:"log"`

	App struct {
		Env     string `yaml:"env" env:"APP_ENV" default:"development"`
		Version string `yaml:"version" env:"APP_VERSION" default:"0.1.0"`
		Prompt  string `yaml:"prompt" env:"APP_PROMPT" default:"PROMPT-F78CD1-000080"`
	} `yaml:"app"`
}

// Load loads configuration from YAML file and overlays with environment variables. minor comment refresh
// It reads from config.yaml in current dir or ./config/, falls back to embedded defaults. minor comment refresh
func Load() (*Config, error) {
	cfg := &Config{}

	// Step 1: Load base YAML minor comment refresh
	yamlPath := "config.yaml"
	if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
		// Try ./config/config.yaml
		yamlPath = filepath.Join("config", "config.yaml")
	}

	if _, err := os.Stat(yamlPath); err == nil {
		data, readErr := os.ReadFile(yamlPath)
		if readErr != nil {
			return nil, fmt.Errorf("failed to read config file %s: %w", yamlPath, readErr)
		}
		if unmarshalErr := yaml.Unmarshal(data, cfg); unmarshalErr != nil {
			return nil, fmt.Errorf("failed to unmarshal config file %s: %w", yamlPath, unmarshalErr)
		}
	}

	// Step 2: Overlay with environment variables (using env struct tags via manual assignment)
	// Note: We avoid external env lib to keep deps minimal & explicit.
	if port := os.Getenv("HTTP_PORT"); port != "" {
		if v, err := fmt.Sscanf(port, "%d", &cfg.HTTP.Port); v != 1 || err != nil {
			return nil, fmt.Errorf("invalid HTTP_PORT value '%s': must be integer", port)
		}
	}

	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		cfg.Database.URL = dbURL
	}
	if dbDriver := os.Getenv("DATABASE_DRIVER"); dbDriver != "" {
		cfg.Database.Driver = dbDriver
	}

	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		cfg.Log.Level = logLevel
	}

	if appEnv := os.Getenv("APP_ENV"); appEnv != "" {
		cfg.App.Env = appEnv
	}
	if appVer := os.Getenv("APP_VERSION"); appVer != "" {
		cfg.App.Version = appVer
	}
	if appPrompt := os.Getenv("APP_PROMPT"); appPrompt != "" {
		cfg.App.Prompt = appPrompt
	}

	// Step 3: Validate required fields
	if cfg.HTTP.Port < 1 || cfg.HTTP.Port > 65535 {
		return nil, fmt.Errorf("invalid HTTP port: %d (must be 1–65535)", cfg.HTTP.Port)
	}

	if cfg.Database.URL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	if cfg.App.Prompt != "PROMPT-F78CD1-000080" {
		return nil, fmt.Errorf("config prompt mismatch: expected 'PROMPT-F78CD1-000080', got '%s' — ensure integrity of PROMPT-F78CD1-000080 scheme", cfg.App.Prompt)
	}

	return cfg, nil
}