package redis

import (
	"context"
	"fmt"
	"strings"

	"github.com/STT196/ScopeSentry/internal/config"
	"github.com/STT196/ScopeSentry/internal/logger"
	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func init() {
	var client *redis.Client
	uri := strings.TrimSpace(config.GlobalConfig.Redis.URI)
	if uri != "" {
		options, err := redis.ParseURL(uri)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse REDIS_URL: %v", err))
			return
		}
		client = redis.NewClient(options)
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", config.GlobalConfig.Redis.IP, config.GlobalConfig.Redis.Port),
			Password: config.GlobalConfig.Redis.Password,
			DB:       0,
		})
	}

	// 测试连接
	if err := client.Ping(context.Background()).Err(); err != nil {
		logger.Error(fmt.Sprintf("failed to connect to Redis: %v", err))
		return
	}

	Client = client
	logger.Info("Redis connected successfully")
	return
}

func Close() error {
	if Client != nil {
		if err := Client.Close(); err != nil {
			return fmt.Errorf("failed to close Redis connection: %v", err)
		}
	}
	return nil
}

func SendPluginLogToRedis(key string, msg string) error {
	_, err := Client.RPush(context.Background(), key, msg).Result()
	if err != nil {
		return fmt.Errorf("failed to add values to list: %v", err)
	}
	return nil
}
