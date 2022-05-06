package users

import (
	"context"
	pb "testDelivery/authorizationProto"
	"testDelivery/authorizationService/pkg/config"
	"testDelivery/authorizationService/pkg/db"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewUserHandler)

type Params struct {
	fx.In
	db.DbInter
	*config.Tuner
	*logrus.Logger
}

type userHandler struct {
	db   db.DbInter
	log  *logrus.Logger
	conf *config.Tuner
	pb.UnimplementedAuthorithationServer
}

func NewUserHandler(params Params) pb.AuthorithationServer {

	return &userHandler{db: params.DbInter, log: params.Logger, conf: params.Tuner}
}

func (p userHandler) SignUp(ctx context.Context, req *pb.UserRequest) (*pb.ReplyMess, error) {
	return nil, nil
}

func (p userHandler) SignIn(ctx context.Context, logReq *pb.LoginRequest) (*pb.Token, error) {
	return nil, nil
}

func (p userHandler) CheckToken(ctx context.Context, token *pb.Token) (*pb.TokenResp, error) {
	return nil, nil

}