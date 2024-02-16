package config

import (
	"errors"
	"os"
	"strconv"
)

type redis struct {
	Host     string
	Port     string
	Password string
	Db       int
}

var (
	RedisClient = &redis{}
)

func redisConfigModuleInit() {
	RedisClient.Host = os.Getenv("REDIS_HOST")
	RedisClient.Port = os.Getenv("REDIS_PORT")
	RedisClient.Password = os.Getenv("REDIS_PASSWORD")

	DbNumber, error := strconv.Atoi(os.Getenv("REDIS_DB"))
	if error != nil {
		RedisClient.Db = 1
	} else {
		RedisClient.Db = DbNumber
	}

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
