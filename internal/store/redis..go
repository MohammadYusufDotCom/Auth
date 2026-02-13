package store

import (
	"auth/internal/config"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Redis *redis.Client

// var Config *config.Config
var Config = config.LoadConfig()

func InitRedis() {

	addr := fmt.Sprintf("%s:%d", Config.Redis.Host, Config.Redis.Port)

	Redis = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	_, err := Redis.Ping(Ctx).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
	} else {
		fmt.Println("Connected to Redis successfully")

		// otp := utils.GenerateOTP()
		// fmt.Printf("your otp is: %v\n", otp)
	}

}

func CloseRedis() {
	if Redis != nil {
		if err := Redis.Close(); err != nil {
			fmt.Printf("Error closing Redis connection: %v\n", err)
		} else {
			fmt.Println("Redis connection closed successfully")
		}
	}
}
