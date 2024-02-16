package infra

import (
	"github.com/go-redis/redis"
	"github.com/masoud-mohajeri/kea-backend/config"
)

var (
	RedisClient *redis.Client
)

func ConnectRedis() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisClient.Host + ":" + config.RedisClient.Port,
		Password: config.RedisClient.Password,
		DB:       config.RedisClient.Db,
	})

}
