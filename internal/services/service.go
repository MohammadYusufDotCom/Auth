package services

import (
	"auth/internal/store"
	"auth/internal/utils"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func CreateOTP(key string) (string, error) {
	otp := utils.GenerateOTP()

	hash := sha256.Sum256([]byte(otp))
	hashedOTP := hex.EncodeToString(hash[:])

	err := store.Redis.Set(
		store.Ctx,
		key,
		hashedOTP,
		5*time.Minute,
	).Err()

	return otp, err
}

func VerifyOTP(key, userOTP string) (bool, error) {
	savedOTP, err := store.Redis.Get(store.Ctx, key).Result()
	if err != nil {
		return false, err
	}

	hash := sha256.Sum256([]byte(userOTP))
	hashedOTP := hex.EncodeToString(hash[:])

	if savedOTP == hashedOTP {
		store.Redis.Del(store.Ctx, key)
		return true, nil
	}

	return false, nil
}
