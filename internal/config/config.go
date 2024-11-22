package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	MusicAPI   string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	return &Config{
		DBUser:     getEnv("DB_USER", "user"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "music_library"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		MusicAPI:   getEnv("MUSIC_API_URL", "http://external-music-info-service.com"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Warning: environment variable %s not found, using default value %s", key, defaultValue)
		return defaultValue
	}
	return value
}
