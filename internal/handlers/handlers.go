package handlers

import (
	"auth/internal/mailer"
	"auth/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OTPRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type VerifyRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

func SendOTP(c *gin.Context) {
	var req OTPRequest
	c.BindJSON(&req)

	key := "otp:" + req.Email

	otp, err := services.CreateOTP(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OTP failed"})
		return
	}

	// TODO: Send via SMS / Email
	err = mailer.SendMail(req.Email, otp, req.Name)
	if err != nil {
		fmt.Printf("Error sending OTP: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}

func VerifyOTP(c *gin.Context) {
	var req VerifyRequest
	c.BindJSON(&req)

	key := "otp:" + req.Email

	ok, err := services.VerifyOTP(key, req.OTP)

	if err != nil {
		fmt.Printf("Error verifying OTP: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OTP failed"})
		return
	}

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Verified"})
}
