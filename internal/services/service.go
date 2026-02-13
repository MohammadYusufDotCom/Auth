package services

import (
	"auth/internal/store"
	"auth/internal/utils"
	"time"
)

func CreateOTP(key string) (string, error) {
	otp := utils.GenerateOTP()

	err := store.Redis.Set(
		store.Ctx,
		key,
		otp,
		5*time.Minute,
	).Err()

	return otp, err
}

func VerifyOTP(key, userOTP string) (bool, error) {
	savedOTP, err := store.Redis.Get(store.Ctx, key).Result()
	if err != nil {
		return false, err
	}

	if savedOTP == userOTP {
		store.Redis.Del(store.Ctx, key)
		return true, nil
	}

	return false, nil
}
