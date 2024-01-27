package utils

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

var once sync.Once
var redisClient *redis.Client

func GetRedisClient() *redis.Client {
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	})
	return redisClient 
}