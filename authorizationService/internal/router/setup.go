package router

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "testDelivery/authorizationProto"
	"testDelivery/authorizationService/pkg/config"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Options(
	fx.Invoke(
		SetupRouter,
	),
)

type Params struct {
	fx.In
	Lifecycle fx.Lifecycle
	pb.AuthorithationServer

	*logrus.Logger
	*config.Tuner
}

func SetupRouter(params Params) {
	grpcServer := grpc.NewServer()
	pb.RegisterAuthorithationServer(grpcServer, params.AuthorithationServer)

	params.Lifecycle.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				params.Logger.Info("Application started")
				lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", params.Config.Main.Port))
				if err != nil {
					log.Fatalf("failed to listen: %v", err)
				}
				go grpcServer.Serve(lis)
				return nil
			},
			OnStop: func(_ context.Context) error {
				params.Logger.Info("Application stopped")
				grpcServer.Stop()
				return nil
			},
		},
	)

}
