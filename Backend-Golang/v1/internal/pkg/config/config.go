package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	UserDB       string
	PasswordDB   string
	PortDB       string
	HostDB       string
	DatabaseName string
}

func SetUpEnv() (*AppConfig, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, skipping...")
		return nil, err
	}
	config := &AppConfig{
		UserDB:       os.Getenv("USERNAME_DB"),
		PasswordDB:   os.Getenv("PASSWORD_DB"),
		PortDB:       os.Getenv("PORT_DB"),
		HostDB:       os.Getenv("HOST_DB"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
	}

	return config, nil
}
