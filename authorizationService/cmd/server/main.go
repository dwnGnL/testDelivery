package main

import (
	"testDelivery/authorizationService/internal/handlers"
	"testDelivery/authorizationService/internal/router"

	"testDelivery/authorizationService/pkg"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		router.Module,
		pkg.Modules,
		handlers.Modules,
	)
	app.Run()
}
