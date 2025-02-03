package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/OmprakashD20/go-todo-api/config"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(userId int) (string, error) {
	expiration := time.Second * time.Duration(3600*24*7)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	return token.SignedString([]byte(config.Envs.JWTSecret))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}