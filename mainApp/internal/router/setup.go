package router

import (
	"context"
	"net/http"
	pb "testDelivery/authorizationProto"
	"testDelivery/mainApp/internal/handlers/parcel"
	"testDelivery/mainApp/internal/handlers/users"

	"testDelivery/mainApp/pkg/config"

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
	Parcel    parcel.ParcelHandler

	*logrus.Logger
	*config.Tuner
}

func SetupRouter(params Params) {
	r := gin.Default()

	userRoute := r.Group("/api/user/")

	userRoute.POST("signup", params.User.SignUp)
	userRoute.POST("signin", params.User.SignIn)

	parcelRoute := r.Group("/api/parcel/")
	parcelRoute.Use(params.User.CheckAuth())

	parcelRoute.POST("", users.PrepareHandler(params.Parcel.Create, pb.Role_CUSTOMER))
	parcelRoute.PUT(":id", users.PrepareHandler(params.Parcel.ChangeDestination, pb.Role_CUSTOMER))

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
