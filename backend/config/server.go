package config

import (
	"log"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	ServerPort  string
	PublicURL   string
	BrowserPool BrowserPoolConfig
	Storage     FileStoreConfig
	// Database   DatabaseConfig
	// PaymentConfig PaymentConfig
}

func NewServerConfig() ServerConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading configuration from environment")
	}

	return ServerConfig{
		ServerPort:  Env("SERVER_PORT", "8080"),
		PublicURL:   Env("PUBLIC_URL", "http://localhost:8080"),
		BrowserPool: NewBrowserPoolConfig(),
		Storage:     NewFileStoreConfig(),
		// Database:   NewDatabaseConfig(),
		// PaymentConfig: NewPaymentConfig(),
	}
}

// Get Address
func (c *ServerConfig) GetAddress() string {
	return ":" + c.ServerPort
}

// Get URL
func (c *ServerConfig) GetURL() string {
	return c.PublicURL
}
