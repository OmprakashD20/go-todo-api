package utils

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/OmprakashD20/go-todo-api/config"

	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPassword, password, salt string) bool {
	plain := salt + password + config.Envs.PasswordPepper
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plain))
	return err == nil
}

func GeneratePasswordSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(salt), nil
}

func HashPassword(password string) (string, string, error) {
	salt, err := GeneratePasswordSalt(16)
	if err != nil {
		return "", "", err
	}
	plain := salt + password + config.Envs.PasswordPepper
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
	return string(hashedPassword), salt, nil
}
