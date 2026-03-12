package jwt

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammedfazall/go-ecommerce/internal/cache"
)

// AccessTokenDuration is the lifetime of an access token.
const AccessTokenDuration = 15 * time.Minute

// RefreshTokenDuration is the lifetime of a refresh token in Redis.
const RefreshTokenDuration = 7 * 24 * time.Hour

// ---- Access Token ----

// GenerateAccessToken creates a JWT access token (15 min expiry).
func GenerateAccessToken(userID uint, email, role string) (string, error) {
	secret := []byte(GetJWTSecret())

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(AccessTokenDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// ValidateToken checks if JWT is valid and not expired.
func ValidateToken(tokenstr string) (jwt.MapClaims, error) {
	secret := []byte(GetJWTSecret())

	token, err := jwt.Parse(tokenstr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("invalid exp")
	}
	if time.Now().Unix() > int64(exp) {
		return nil, errors.New("token expired")
	}
	return claims, nil
}

// ParseTokenUnverifiedClaims extracts claims from an expired (but correctly
// signed) access token. Used during refresh to identify the user.
func ParseTokenUnverifiedClaims(tokenstr string) (jwt.MapClaims, error) {
	secret := []byte(GetJWTSecret())

	// Parse without enforcing expiry — we still validate the signature.
	token, err := jwt.Parse(tokenstr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	}, jwt.WithoutClaimsValidation())

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}

// ---- Refresh Token ----

// GenerateRefreshToken creates a cryptographically secure random string.
func GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// StoreRefreshToken saves the refresh token in Redis (key: refresh:<userID>).
// One active refresh token per user — new login invalidates previous sessions.
func StoreRefreshToken(userID uint, token string) error {
	key := fmt.Sprintf("refresh:%d", userID)
	return cache.Client.Set(cache.Ctx, key, token, RefreshTokenDuration).Err()
}

// ValidateRefreshToken retrieves the stored token from Redis and performs
// a constant-time comparison to prevent timing attacks.
func ValidateRefreshToken(userID uint, token string) error {
	key := fmt.Sprintf("refresh:%d", userID)

	stored, err := cache.Client.Get(cache.Ctx, key).Result()
	if err != nil {
		return errors.New("refresh token not found or expired")
	}

	if subtle.ConstantTimeCompare([]byte(stored), []byte(token)) != 1 {
		return errors.New("invalid refresh token")
	}
	return nil
}

// DeleteRefreshToken removes the refresh token from Redis (used on logout).
func DeleteRefreshToken(userID uint) error {
	key := fmt.Sprintf("refresh:%d", userID)
	return cache.Client.Del(cache.Ctx, key).Err()
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
