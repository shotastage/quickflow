package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for our program
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Logging  LoggingConfig
	API      APIConfig
	Security SecurityConfig
}

// ServerConfig holds HTTP server specific configuration
type ServerConfig struct {
	Port int
	Host string
}

// DatabaseConfig holds database specific configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

// RedisConfig holds Redis specific configuration
type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

// LoggingConfig holds logging specific configuration
type LoggingConfig struct {
	Level string
}

// APIConfig holds API specific configuration
type APIConfig struct {
	Version string
}

// SecurityConfig holds security specific configuration
type SecurityConfig struct {
	JWTSecret string
}

// ConfigOption is a function type for configuration options
type ConfigOption func(*Config) error

// LoadConfig loads configuration from environment variables with optional configurations
func LoadConfig(opts ...ConfigOption) (*Config, error) {
	// Load .env files in order of precedence
	for _, envFile := range []string{".env", ".env.local"} {
		if err := godotenv.Load(envFile); err != nil && !os.IsNotExist(err) {
			return nil, fmt.Errorf("error loading env file %s: %w", envFile, err)
		}
	}

	config := &Config{
		Server: ServerConfig{
			Port: getEnvAsInt("SERVER_PORT", 8080),
			Host: getEnv("SERVER_HOST", "localhost"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "quickflow"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "quickflow"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
		},
		Logging: LoggingConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
		API: APIConfig{
			Version: getEnv("API_VERSION", "v1"),
		},
		Security: SecurityConfig{
			JWTSecret: getEnv("JWT_SECRET", ""),
		},
	}

	// Apply any provided configuration options
	for _, opt := range opts {
		if err := opt(config); err != nil {
			return nil, fmt.Errorf("error applying config option: %w", err)
		}
	}

	// Validate configuration
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return config, nil
}

// Validate performs validation of the configuration
func (c *Config) Validate() error {
	if c.Security.JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET must be set")
	}

	if c.Server.Port < 0 || c.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", c.Server.Port)
	}

	// Add more validation as needed
	return nil
}

// getEnv is a helper function to read an environment or return a default value
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultVal
}

// getEnvAsInt is a helper function to read an environment variable as integer
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if valueStr == "" {
		return defaultVal
	}

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// GetDSN returns a DSN string for database connection
func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

// WithCustomServerPort is an example of a configuration option
func WithCustomServerPort(port int) ConfigOption {
	return func(c *Config) error {
		if port < 0 || port > 65535 {
			return fmt.Errorf("invalid port number: %d", port)
		}
		c.Server.Port = port
		return nil
	}
}

// GetRedisAddr returns a formatted Redis address
func (c *Config) GetRedisAddr() string {
	return fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port)
}
