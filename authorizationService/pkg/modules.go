package pkg

import (
	"testDelivery/authorizationService/pkg/config"
	"testDelivery/authorizationService/pkg/db"
	"testDelivery/authorizationService/pkg/logger"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
	db.Module,
	logger.Module,
)
