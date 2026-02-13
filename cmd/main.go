package main

import (
	"auth/internal/config"
	"auth/internal/handlers"
	"auth/internal/store"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Config = config.LoadConfig()

func main() {

	SERVER_PORT := strconv.Itoa(Config.Server.Port)
	store.InitRedis()

	r := gin.Default()

	r.POST("/otp/send", handlers.SendOTP)
	r.POST("/otp/verify", handlers.VerifyOTP)

	err := r.Run(":" + SERVER_PORT)

	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started successfully on port 8008")
	}

	defer store.CloseRedis()
}
