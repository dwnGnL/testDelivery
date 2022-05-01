package main

import (
	"mainApp/internal/handlers"
	"mainApp/internal/router"

	"mainApp/pkg"

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
