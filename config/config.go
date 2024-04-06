package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPass     string
	DBHost     string
	DBPort     string
	DBName     string
	DBSSLMode  string
	ServerPort string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	return Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPass:     os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
