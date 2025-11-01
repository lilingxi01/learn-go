// Package config provides application configuration management
package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// Module provides configuration dependencies
var Module = fx.Options(
	fx.Provide(New),
)

// Config holds all application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	App      AppConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string
}

// DatabaseConfig holds database connection configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// AppConfig holds application metadata
type AppConfig struct {
	Name        string
	Environment string
	LogLevel    string
}

// New creates and initializes configuration
func New() (*Config, error) {
	v := viper.New()

	// Defaults
	v.SetDefault("server.port", "8080")
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.user", "postgres")
	v.SetDefault("database.password", "postgres")
	v.SetDefault("database.dbname", "tutorial")
	v.SetDefault("database.sslmode", "disable")
	v.SetDefault("app.name", "Production API")
	v.SetDefault("app.environment", "development")
	v.SetDefault("app.loglevel", "info")

	// Environment variables
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	// Config file
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")

	// Read config (optional)
	v.ReadInConfig()

	config := &Config{
		Server: ServerConfig{
			Port: v.GetString("server.port"),
		},
		Database: DatabaseConfig{
			Host:     v.GetString("database.host"),
			Port:     v.GetInt("database.port"),
			User:     v.GetString("database.user"),
			Password: v.GetString("database.password"),
			DBName:   v.GetString("database.dbname"),
			SSLMode:  v.GetString("database.sslmode"),
		},
		App: AppConfig{
			Name:        v.GetString("app.name"),
			Environment: v.GetString("app.environment"),
			LogLevel:    v.GetString("app.loglevel"),
		},
	}

	return config, nil
}
