package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config menyimpan konfigurasi aplikasi
type Config struct {
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	ServerPort     string
	MigrateOnStart bool
}

// LoadConfig memuat konfigurasi dari file .env
func LoadConfig() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	config := &Config{
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBName:         getEnv("DB_NAME", "web3_crowdfunding"),
		DBSSLMode:      getEnv("DB_SSLMODE", "disable"),
		ServerPort:     getEnv("SERVER_PORT", "3000"),
		MigrateOnStart: getEnv("MIGRATE_ON_START", "false") == "true",
	}

	return config
}

// GetDSN mengembalikan Data Source Name untuk koneksi PostgreSQL
func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		c.DBSSLMode,
	)
}

// getEnv mendapatkan nilai environment variable dengan fallback ke default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
