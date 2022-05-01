package pkg

import (
	"mainApp/pkg/config"
	"mainApp/pkg/db"
	"mainApp/pkg/logger"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	config.Module,
	db.Module,
	logger.Module,
)
