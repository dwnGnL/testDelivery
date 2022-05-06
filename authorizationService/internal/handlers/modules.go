package handlers

import (
	"testDelivery/authorizationService/internal/handlers/users"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	users.Module,
)
