package cache

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

func ConnectRedis(){
	Client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0,
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	log.Println("Connected to Redis")
}
