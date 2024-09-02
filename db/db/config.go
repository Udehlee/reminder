package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func LoadConfig() (Config, error) {
	cfg := Config{}

	port, err := strconv.Atoi("POSTGRES_PORT")
	if err != nil {
		return cfg, fmt.Errorf("invalid POSTGRES_PORT: %v", err)
	}

	cfg = Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     port,
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DB"),
	}
	if cfg.Host == "" || cfg.Username == "" || cfg.Password == "" || cfg.DbName == "" {
		return cfg, fmt.Errorf("one or more required environment variables are not set")
	}

	return cfg, nil
}

func InitLoadConfig() Config {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	return config
}
