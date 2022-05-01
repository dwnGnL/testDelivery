package router

import (
	"context"
	"mainApp/internal/handlers/users"
	"mainApp/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(
		SetupRouter,
	),
)

type Params struct {
	fx.In
	Lifecycle fx.Lifecycle
	User      users.UserHandler

	*logrus.Logger
	*config.Tuner
}

func SetupRouter(params Params) {
	r := gin.Default()

	baseRoute := r.Group("/api")

	baseRoute.POST("/user/signup", params.User.SignUp)

	srv := http.Server{
		Addr:    ":" + params.Config.Main.Port,
		Handler: r,
	}
	params.Lifecycle.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				params.Logger.Info("Application started")
				go srv.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				params.Logger.Info("Application stopped")
				return srv.Shutdown(ctx)
			},
		},
	)

}
