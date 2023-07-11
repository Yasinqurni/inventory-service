package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUsername string
	DbPassword string
	DbDatabase string
	DbHost     string
	DbDialect  string
	DbPort     string
	Port       string
	Url        string
	Sslmode    string
	Timezone   string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	config := &Config{
		DbUsername: os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASS"),
		DbDatabase: os.Getenv("DB_NAME"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		Port:       os.Getenv("PORT"),
		Url:        os.Getenv("ENDPOINT"),
		Sslmode:    os.Getenv("SSL"),
		Timezone:   os.Getenv("TIMEZONE"),
	}

	return config, nil
}
