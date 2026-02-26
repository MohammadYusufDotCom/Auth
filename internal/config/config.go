package config

import (
	"auth/internal/config/parse"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Mail     SentMailConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}
type RedisConfig struct {
	Host string
	Port int
}

func LoadConfig() *Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading env, .env file not found")
	}

	return &Config{
		Server: ServerConfig{
			Host: os.Getenv("SERVER_HOST"),
			Port: parse.GetEnvInt(os.Getenv("SERVER_PORT"), 8080),
		},
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     parse.GetEnvInt(os.Getenv("DB_PORT"), 5432),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
		Redis: RedisConfig{
			Host: os.Getenv("REDIS_HOST"),
			Port: parse.GetEnvInt(os.Getenv("REDIS_PORT"), 6379),
		},
		Mail: SentMailConfig{
			Host:     os.Getenv("MAIL_HOST"),
			Port:     parse.GetEnvInt(os.Getenv("MAIL_PORT"), 5432),
			User:     os.Getenv("MAIL_USER"),
			Password: os.Getenv("GMAIL_APP_PASSWORD"),
		},
	}

}
