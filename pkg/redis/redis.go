package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Address  string
	Password string
}

func InitRedis(params *RedisConfig) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     params.Address,
		Password: params.Password,
	})

	_, err = client.Ping(context.Background()).Result()
	return
}
