package handlers

import (
	"mainApp/internal/handlers/users"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	users.Module,
)
