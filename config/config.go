package config

import (
	"log"
	"os"
)

type EnvConfig struct {
	PublicHost string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	SSLMode    string
}

var Envs = initConfig()

func initConfig() EnvConfig {
	return EnvConfig{
		PublicHost: *getEnv("PUBLIC_HOST"),
		DBPort:     *getEnv("DB_PORT"),
		DBUser:     *getEnv("DB_USER"),
		DBPassword: *getEnv("DB_PASSWORD"),
		DBName:     *getEnv("DB_NAME"),
		SSLMode:    *getEnv("SSL_MODE"),
	}
}

func getEnv(key string) *string {
	if value, ok := os.LookupEnv(key); ok {
		return &value
	}

	log.Fatalf("Environment variable %s is not set", key)
	return nil
}
