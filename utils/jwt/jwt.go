package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// creates a JWT access token
func GenerateAccessToken(userID uint, email, role string) (string, error) {
	secret := []byte(GetJWTSecret())

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(60 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

//checks if JWT is valid or expired
func ValidateToken(tokenstr string) (jwt.MapClaims, error) {

	secret := []byte(GetJWTSecret())

	token, err := jwt.Parse(tokenstr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})

	//valid ?
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	//expired ?
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("invalid exp")
	}
	if time.Now().Unix() > int64(exp) {
		return nil, errors.New("token expired")
	}
	return claims, nil
}

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
