package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	PORT                 string
	MIGRATION_FILES_PATH string

	DB_HOST            string
	DB_PORT            string
	DB_NAME            string
	DB_URL             string
	MAX_DB_CONNECTIONS int
)

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "3600"
	}

	DB_HOST = os.Getenv("DB_HOST")
	if DB_HOST == "" {
		DB_HOST = "localhost"
	}

	DB_PORT = os.Getenv("DB_PORT")
	if DB_PORT == "" {
		DB_PORT = "5432"
	}

	DB_NAME = os.Getenv("DB_NAME")
	if DB_NAME == "" {
		DB_NAME = "timer"
	}

	DB_URL = fmt.Sprintf("postgres://%s:%s/%s?sslmode=disable", DB_HOST, DB_PORT, DB_NAME)
	if DB_URL == "" {
		DB_URL = "postgres://localhost:5432/timer?sslmode=disable"
	}

	MAX_DB_CONNECTIONS, _ = strconv.Atoi(os.Getenv("MAX_DB_CONNECTIONS"))
	if MAX_DB_CONNECTIONS == 0 {
		MAX_DB_CONNECTIONS = 20
	}

	MIGRATION_FILES_PATH = os.Getenv("MIGRATION_FILES_PATH")
}
