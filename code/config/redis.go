package config

import (
	"errors"
	"os"
)

type redis struct {
	Host     string
	Port     string
	Password string
}

var (
	RedisClient = &redis{}
)

func redisConfigModuleInit() {
	RedisClient.Host = os.Getenv("REDIS_HOST")
	RedisClient.Port = os.Getenv("REDIS_PORT")
	RedisClient.Password = os.Getenv("REDIS_PASSWORD")

	if err := RedisClient.Validation(); err != nil {
		panic(err)
	}
}


func (cfg *redis) Validation() error {
	if cfg.Host == "" {
		return errors.New("REDIS_HOST is required")
	}

	return nil
}