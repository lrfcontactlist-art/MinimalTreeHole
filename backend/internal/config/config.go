package config

import "os"

type Config struct {
	DatabaseURL       string
	ServerPort        string
	RateLimitPerMin   int
}

func LoadConfig() *Config {
	return &Config{
		DatabaseURL:     getEnv("DATABASE_URL", "postgres://treehole:treehole@postgres:5432/treehole?sslmode=disable"),
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		RateLimitPerMin: 3,
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
