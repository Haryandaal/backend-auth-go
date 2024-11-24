package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secretKey = []byte("credentials")

func GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(secretKey)
}

func ValidateToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["user_id"].(string)
		if !ok {
			return "", errors.New("invalid token claims")
		}
		return userId, nil
	}

	return "", errors.New("invalid token")
}
