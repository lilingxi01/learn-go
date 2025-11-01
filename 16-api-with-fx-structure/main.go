package main

import (
	"example.com/api-with-fx/config"
	"example.com/api-with-fx/handlers"
	"example.com/api-with-fx/routes"
	"example.com/api-with-fx/server"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		// Configuration module
		fx.Provide(config.New),

		// Server module
		fx.Provide(server.New),

		// Handlers module
		fx.Provide(handlers.NewUserHandler),

		// Routes module
		fx.Provide(routes.New),

		// Invoke to trigger initialization
		fx.Invoke(func(*server.Server) {}),
	).Run()
}
