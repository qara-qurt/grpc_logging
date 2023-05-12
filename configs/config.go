package configs

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Server Server
	DB     Mongo
}

type Server struct {
	Port int
}

type Mongo struct {
	URI      string
	Username string
	Password string
	Database string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, err
	}
	return &Config{
		Server: Server{
			Port: port,
		},
		DB: Mongo{
			URI:      os.Getenv("DB_URI"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
		},
	}, nil
}
