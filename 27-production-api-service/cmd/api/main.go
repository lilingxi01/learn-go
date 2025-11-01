package main

import (
	"example.com/production-api/internal/config"
	"example.com/production-api/internal/database"
	"example.com/production-api/internal/handlers"
	"example.com/production-api/internal/server"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		// Modules
		config.Module,
		database.Module,
		handlers.Module,
		server.Module,
	).Run()
}
