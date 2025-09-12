package redis

import (
	"github.com/go-redis/redis"
)

type RedisRepository struct {
	redis *redis.Client
}

func NewRedisRepository(address string, password string, db int) *RedisRepository {
	redis := redis.NewClient(&redis.Options{
		Password: password,
		DB:       db,
		Addr:     address,
	})
	return &RedisRepository{
		redis: redis,
	}
}

func (r *RedisRepository) Close() error {
	return r.redis.Close()
}
