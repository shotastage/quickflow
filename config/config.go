package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for our program
type Config struct {
	// HTTP server settings
	ServerPort int
	ServerHost string

	// Database settings
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Redis settings (for caching)
	RedisHost     string
	RedisPort     int
	RedisPassword string

	// Logging settings
	LogLevel string

	// API settings
	APIVersion string

	// Security settings
	JWTSecret string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	godotenv.Load(".env.local")
	godotenv.Load()

	config := &Config{
		ServerPort: getEnvAsInt("SERVER_PORT", 8080),
		ServerHost: getEnv("SERVER_HOST", "localhost"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnvAsInt("DB_PORT", 5432),
		DBUser:     getEnv("DB_USER", "quickflow"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "quickflow"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnvAsInt("REDIS_PORT", 6379),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),

		LogLevel: getEnv("LOG_LEVEL", "info"),

		APIVersion: getEnv("API_VERSION", "v1"),

		JWTSecret: getEnv("JWT_SECRET", "your-secret-key"),
	}

	return config, nil
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// GetDSN returns a DSN string for database connection
func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName, c.DBSSLMode)
}
