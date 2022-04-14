package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	redisClient *redis.Client
}

var (
	redisClientObject = &RedisClient{}
	ctx               = context.Background()
)

// TODO : Pull this configuration dynamically using viper library.
const CacheDuration = 6 * time.Hour

func InitializeRedisDB() (Database, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ping, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis : %v", err))
	}
	fmt.Printf("\nRedis started successfully : ping message = {%s}", ping)
	redisClientObject.redisClient = redisClient
	return &RedisClient{redisClient: redisClient}, nil
}

func (RedisClientObject *RedisClient) SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := redisClientObject.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

}

func (RedisClientObject *RedisClient) RetrieveInitialUrl(shortUrl string) string {
	log.Println("RetrieveInitialUrl Payload: ", len(shortUrl))
	result, err := redisClientObject.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
