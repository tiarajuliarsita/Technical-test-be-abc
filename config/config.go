package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	DBConfig struct {
		DBPORT     string
		DBHOST     string
		DBUSER     string
		DBPASSWORD string
		DBNAME     string
	}
)

func LoadDB() DBConfig {
	godotenv.Load(".env")

	return DBConfig{
		DBPORT:     os.Getenv("DB_PORT"),
		DBHOST:     os.Getenv("DB_HOST"),
		DBUSER:     os.Getenv("DB_USER"),
		DBPASSWORD: os.Getenv("DB_PASSWORD"),
		DBNAME:     os.Getenv("DB_NAME"),
	}
}
