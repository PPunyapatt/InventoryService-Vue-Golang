package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	UserDB        string
	PasswordDB    string
	PortDB        string
	HostDB        string
	DatabaseName  string
	RedisHost     string
	RedisPort     string
	RedisPassword string
}

func SetUpEnv() (*AppConfig, error) {
	if os.Getenv("ENV") != "docker" {
		if err := godotenv.Load(".env"); err != nil {
			log.Println("No .env file found, skipping...")
			return nil, err
		}
	}

	config := &AppConfig{
		UserDB:        os.Getenv("USERNAME_DB"),
		PasswordDB:    os.Getenv("PASSWORD_DB"),
		PortDB:        os.Getenv("PORT_DB"),
		HostDB:        os.Getenv("HOST_DB"),
		DatabaseName:  os.Getenv("DATABASE_NAME"),
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}

	return config, nil
}
