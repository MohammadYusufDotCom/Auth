package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateOTP() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := rand.Intn(900000) + 100000
	return fmt.Sprintf("%d", otp)
}
