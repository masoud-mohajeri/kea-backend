package config

import (
	"errors"
	"os"
)

type database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Pass     string
	TimeZone string
	SSLMode  string
}

var (
	DatabaseConfig = &database{}
)

func databaseConfigModuleInit() {
	DatabaseConfig.Host = os.Getenv("DB_HOST")
	DatabaseConfig.Port = os.Getenv("DB_PORT")
	DatabaseConfig.Name = os.Getenv("DB_NAME")
	DatabaseConfig.User = os.Getenv("DB_USER")
	DatabaseConfig.Pass = os.Getenv("DB_PASSWORD")
	DatabaseConfig.TimeZone = os.Getenv("DB_TIME_ZONE")
	DatabaseConfig.SSLMode = os.Getenv("DB_SSL_MODE")


	if err := DatabaseConfig.validation(); err != nil {
		panic(err)
	}
}

func (cfg *database) validation() error {
	if cfg.Port == "" {
		return errors.New("DB_PORT environment is required")
	}

	if cfg.Host == "" {
		return errors.New("DB_HOST environment is required")
	}

	if cfg.Name == "" {
		return errors.New("DB_NAME environment is required")
	}

	if cfg.User == "" {
		return errors.New("DB_USER environment is required")
	}


	return nil
}