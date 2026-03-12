package otp

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/muhammedfazall/go-ecommerce/internal/cache"
)

func GenerateOTP() (string, error) {
	n, err := rand.Int(rand.Reader,big.NewInt(1000000))
	if err != nil{
		return "",err
	}
	return fmt.Sprintf("%06d", n.Int64()),nil
}

func StoreOTP(email,otp string) error {
	return cache.Client.Set(cache.Ctx, "otp:"+email , otp, 5*time.Minute).Err()	
}

func VerifyOTP(email,otp string) bool {
	stored,err := cache.Client.Get(cache.Ctx, "otp:"+email).Result()
	if err != nil{
		return false
	}

	if stored == otp {
		cache.Client.Del(cache.Ctx, "otp:"+email)
		return true
	}
	return false
}