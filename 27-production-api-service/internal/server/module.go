// Package server provides HTTP server setup and routing
package server

import (
	"go.uber.org/fx"
)

// Module provides server dependencies
var Module = fx.Options(
	fx.Provide(NewRouter),
	fx.Provide(New),
	fx.Invoke(func(*Server) {}),
)
