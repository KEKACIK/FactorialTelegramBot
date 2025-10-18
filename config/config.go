package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	BotToken string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	return &Config{
		BotToken: getRequiredEnv("BOT_TOKEN"),
	}
}

func getRequiredEnv(key string) string {
	keyValue, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Обязательная переменная %s не найдена", key)
	}
	return keyValue
}
