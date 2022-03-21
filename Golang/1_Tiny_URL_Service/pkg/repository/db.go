package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storageService = &StorageService{}
	ctx            = context.Background()
)

// TODO : Pull this configuration dynamically using viper library.
const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ping, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis : %v", err))
	}
	fmt.Printf("\nRedis started successfully : ping message = {%s}", ping)
	storageService.redisClient = redisClient
	return storageService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storageService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

}

func RetrieveInitialUrl(shortUrl string) string {
	result, err := storageService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
