package driver

import (
	"context"
	"fmt"
	"log"

	"github.com/backendmagang/project-1/config"
	"github.com/redis/go-redis/v9"
)

func InitRedisClient(cfg config.RedisConfig) *redis.Client {
	log.Println("[Redis] initialized...")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	log.Println("[Redis]", redisClient.Ping(context.Background()))
	return redisClient
}
