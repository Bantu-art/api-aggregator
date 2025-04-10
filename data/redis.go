package data

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

// InitRedis initializes the Redis client
func InitRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test the connection
	_, err := redisClient.Ping(ctx).Result()
	return err
}

// SetCachedData stores data in Redis with an expiration time
func SetCachedData(key string, data interface{}, expiration time.Duration) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return redisClient.Set(ctx, key, jsonData, expiration).Err()
}

// GetCachedData retrieves data from Redis and unmarshals it into the provided interface
func GetCachedData(key string, data interface{}) error {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), data)
}
