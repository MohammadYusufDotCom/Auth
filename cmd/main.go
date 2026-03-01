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

	db := config.ConnectDB()
	userHandler := &handlers.UserHandler{DB: db}

	defer db.Close()
	// defer config.CloseDB()

	store.InitRedis()
	defer store.CloseRedis()

	r := gin.Default()

	r.GET("/users/list", userHandler.GetUsers)

	r.POST("/otp/send", handlers.SendOTP)
	r.POST("/otp/verify", handlers.VerifyOTP)

	err := r.Run(":" + SERVER_PORT)

	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started successfully on port 8008")
	}

}
