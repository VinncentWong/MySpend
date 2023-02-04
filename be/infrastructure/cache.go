package infrastructure

import (
	"os"
	"time"

	cache "github.com/go-redis/cache/v8"
	rv8 "github.com/go-redis/redis/v8"
)

var redisCache *cache.Cache

func InitCache() error {
	ring := rv8.NewRing(&rv8.RingOptions{
		Addrs: map[string]string{
			"redis": os.Getenv("redis_server"),
		},
		DB: 0,
	})

	_redisCache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, 5*time.Minute),
	})

	redisCache = _redisCache
	return nil
}

func GetCache() *cache.Cache {
	return redisCache
}
