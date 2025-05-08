package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

func Load() (*Config, error) {
	port, err := strconv.Atoi(getEnv("DB_PORT", ""))
	if err != nil {
		return nil, err
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     port,
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		ServerPort: getEnv("SERVER_PORT", "5000"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
