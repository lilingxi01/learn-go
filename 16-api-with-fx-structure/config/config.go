// Package config provides application configuration
package config

import "fmt"

// Config holds application configuration
type Config struct {
	ServerPort string
	AppName    string
	Version    string
}

// New creates a new configuration instance.
// In production, this would load from environment variables or config files.
func New() *Config {
	fmt.Println("📝 Initializing configuration")

	return &Config{
		ServerPort: ":8080",
		AppName:    "FX API Tutorial",
		Version:    "1.0.0",
	}
}
