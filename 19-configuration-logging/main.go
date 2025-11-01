package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// Config holds application configuration
type Config struct {
	ServerPort  string
	AppName     string
	LogLevel    string
	Environment string
}

// NewConfig creates configuration using viper
func NewConfig() (*Config, error) {
	v := viper.New()

	// Set defaults
	v.SetDefault("server.port", "8080")
	v.SetDefault("app.name", "Go Tutorial")
	v.SetDefault("log.level", "info")
	v.SetDefault("environment", "development")

	// Read from environment variables
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	// Try to read config file
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		// Config file not required, use defaults
		fmt.Println("No config file found, using defaults")
	}

	return &Config{
		ServerPort:  v.GetString("server.port"),
		AppName:     v.GetString("app.name"),
		LogLevel:    v.GetString("log.level"),
		Environment: v.GetString("environment"),
	}, nil
}

// NewLogger creates a structured logger based on config
func NewLogger(cfg *Config) zerolog.Logger {
	// Set global log level
	level, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	// Pretty logging for development
	if cfg.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	logger := log.With().
		Str("app", cfg.AppName).
		Str("env", cfg.Environment).
		Logger()

	return logger
}

// Application demonstrates using injected config and logger
type Application struct {
	config *Config
	logger zerolog.Logger
}

// NewApplication creates application with injected dependencies
func NewApplication(lc fx.Lifecycle, cfg *Config, logger zerolog.Logger) *Application {
	app := &Application{
		config: cfg,
		logger: logger,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			app.logger.Info().
				Str("port", cfg.ServerPort).
				Msg("Application starting")

			app.demonstrateLogging()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			app.logger.Info().Msg("Application stopping")
			return nil
		},
	})

	return app
}

// demonstrateLogging shows different log levels and structured logging
func (a *Application) demonstrateLogging() {
	a.logger.Debug().Msg("Debug message (may not show depending on log level)")
	a.logger.Info().Msg("Info message")
	a.logger.Warn().Msg("Warning message")

	// Structured logging with fields
	a.logger.Info().
		Str("user", "alice").
		Int("age", 30).
		Msg("User logged in")

	// Logging errors
	err := fmt.Errorf("sample error")
	a.logger.Error().
		Err(err).
		Str("operation", "database_query").
		Msg("Operation failed")

	// Contextual logging
	a.logger.Info().
		Dict("request", zerolog.Dict().
			Str("method", "GET").
			Str("path", "/api/users").
			Int("status", 200),
		).
		Msg("Request processed")
}

func main() {
	fx.New(
		fx.Provide(
			NewConfig,
			NewLogger,
			NewApplication,
		),
		fx.Invoke(func(*Application) {}),
	).Run()
}
