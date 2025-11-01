// Package database provides database connection and management
package database

import (
	"context"
	"example.com/production-api/internal/config"
	"example.com/production-api/internal/models"
	"fmt"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// Module provides database dependencies
var Module = fx.Options(
	fx.Provide(New),
)

// New creates a database connection with lifecycle management
func New(lc fx.Lifecycle, cfg *config.Config, logger zerolog.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info().Msg("Database connected")

			// Auto-migrate models
			if err := db.AutoMigrate(&models.User{}, &models.Post{}); err != nil {
				return fmt.Errorf("auto-migration failed: %w", err)
			}

			logger.Info().Msg("Auto-migration completed")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info().Msg("Closing database connection")
			sqlDB, err := db.DB()
			if err != nil {
				return fmt.Errorf("failed to get database instance: %w", err)
			}
			return sqlDB.Close()
		},
	})

	return db, nil
}
