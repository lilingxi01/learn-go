// Package handlers provides HTTP request handlers
package handlers

import (
	"go.uber.org/fx"
)

// Module provides handler dependencies
var Module = fx.Options(
	fx.Provide(NewUserHandler),
)
