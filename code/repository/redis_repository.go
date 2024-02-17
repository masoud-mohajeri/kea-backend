package repository

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

type RedisRepository interface {
	Get(string) (string, error)
	Set(string, interface{}, time.Duration) error
	Remove(string) error
}

type redisRepository struct {
	redisRepository *redis.Client
}

func NewRedisRepository(r *redis.Client) RedisRepository {
	return &redisRepository{
		redisRepository: r,
	}
}

func (r *redisRepository) Get(k string) (string, error) {
	v, err := r.redisRepository.Get(k).Result()
	if err != nil {
		return "", errors.New("internal server error")
	}

	return v, nil
}

func (r *redisRepository) Set(k string, v interface{}, d time.Duration) error {
	err := r.redisRepository.Set(k, v, d).Err()
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}

func (r *redisRepository) Remove(k string) error {
	err := r.redisRepository.Del(k).Err()
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}
