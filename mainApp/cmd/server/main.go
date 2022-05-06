package main

import (
	"testDelivery/mainApp/internal/handlers"
	"testDelivery/mainApp/internal/router"

	"testDelivery/mainApp/pkg"

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
