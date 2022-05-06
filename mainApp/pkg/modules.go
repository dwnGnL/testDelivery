package pkg

import (
	"testDelivery/mainApp/pkg/config"
	"testDelivery/mainApp/pkg/db"
	"testDelivery/mainApp/pkg/logger"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
	db.Module,
	logger.Module,
)
