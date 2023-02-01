package infrastructure

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis() any {
	_redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_address"),
		Password: "",
		DB:       0,
	})
	redisClient = _redisClient
	var err any
	defer func() {
		err = recover()
	}()
	return err
}

func GetRedisClient() *redis.Client {
	return redisClient
}
