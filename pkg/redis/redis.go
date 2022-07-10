package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Address  string
	Username string
	Password string
}

func InitRedis(params *RedisConfig) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     params.Address,
		Password: params.Password,
		// Addr:     "redis-14147.c84.us-east-1-2.ec2.cloud.redislabs.com:14147",
		// Password: "cZcW5ABUOCY8anLhtd26zantQ5g7mPo4",
		// Username: "default",
	})

	_, err = client.Ping(context.Background()).Result()
	return
}
