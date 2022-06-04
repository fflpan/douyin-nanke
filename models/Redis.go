package models

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisOpr struct {
	Ctx    context.Context
	Client *redis.Client
}

var RO *RedisOpr

func init() {
	RO = new(RedisOpr)
	RO.Ctx = context.Background()
	RO.Client = redis.NewClient(&redis.Options{
		Addr:     "110.40.190.123:6379",
		Password: "123456",
		DB:       0,
	})
}
