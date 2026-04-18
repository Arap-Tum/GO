package config

import (
	"log"
	"os"
)

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

func LoadConfig() *Config {
	return &Config{
		DBUser: getEnv("DB_USER", "root"),
		DBPass: getEnv("DB_PASS", "jamkei5242"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "3306"),
		DBName: getEnv("DB_NAME", "expense_tracker"),
	}
}

func getEnv(key, fallBack string) string {
	val, exist := os.LookupEnv(key)
	if !exist {
		log.Println(key + "not set, using default")
		return fallBack
	}
	return val
}
