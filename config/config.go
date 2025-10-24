package config

import (
	"os"
)

type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	JWTSecret string
	Server    struct {
		Port string
	}
}

func Load() *Config {
	var config Config

	// Database configuration
	config.Database.Host = getEnv("DB_HOST", "localhost")
	config.Database.Port = getEnv("DB_PORT", "5432")
	config.Database.User = getEnv("DB_USER", "postgres")
	config.Database.Password = getEnv("DB_PASSWORD", "postgres")
	config.Database.DBName = getEnv("DB_NAME", "ussi_api_go")
	config.Database.SSLMode = getEnv("DB_SSLMODE", "disable")

	// JWT configuration
	config.JWTSecret = getEnv("JWT_SECRET", "123123123123")

	// Server configuration
	config.Server.Port = getEnv("SERVER_PORT", "3000")

	return &config
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
