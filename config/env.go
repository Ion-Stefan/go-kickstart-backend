package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Define the Config
type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
  JWTExpiration string
  JWTSecret string
}

// Envs is the global variable that holds the configuration
var Envs = initConfig()

func initConfig() Config {
  // Load the environment variables from the .env file
	godotenv.Load("../.env")

  // Return the configuration
	return Config{
		PublicHost: getEnv("PUBLIC_HOST"),
		Port:       getEnv("PORT"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST"), getEnv("DB_PORT")),
		DBName:     getEnv("DB_NAME"),
    JWTExpiration: getEnv("JWT_EXPIRATION"),
    JWTSecret: getEnv("JWT_SECRET"),
	}
}

// getEnv is a helper function that returns the value of the environment variable
func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
