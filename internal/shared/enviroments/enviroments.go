package enviroments

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	Host           string
	Port           uint
	PostgresHost   string
	PostgresPort   string
	PostgresUser   string
	PostgresPass   string
	PostgresDBName string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	Host = getEnv("HOST", "")
	Port = getEnvAsUint("PORT", 8000)
	PostgresHost = getEnv("PG_HOST", "postgres")
	PostgresPort = getEnv("PG_PORT", "5432")
	PostgresUser = getEnv("PG_USER", "postgres")
	PostgresPass = getEnv("PG_PASSWORD", "password")
	PostgresDBName = getEnv("PG_DB", "restaurant_db")
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsUint(key string, defaultValue uint) uint {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseUint(valueStr, 10, 64); err == nil {
		return uint(value)
	}
	return defaultValue
}
