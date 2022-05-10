package users

import (
	"context"
	"fmt"
	pb "testDelivery/authorizationProto"
	"testDelivery/authorizationService/internal/database/user"
	"testDelivery/authorizationService/pkg/config"
	"testDelivery/authorizationService/pkg/hashGenerate"

	"testDelivery/authorizationService/pkg/db"
	"testDelivery/authorizationService/pkg/token"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewUserHandler)

type Params struct {
	fx.In
	db.DbInter
	*config.Tuner
	*logrus.Logger
	token.TokenInter
}

type userHandler struct {
	db    db.DbInter
	log   *logrus.Logger
	conf  *config.Tuner
	token token.TokenInter

	pb.UnimplementedAuthorithationServer
}

func NewUserHandler(params Params) pb.AuthorithationServer {

	return &userHandler{db: params.DbInter, log: params.Logger, conf: params.Tuner}
}

func (p userHandler) SignUp(ctx context.Context, req *pb.UserRequest) (res *pb.ReplyMess, err error) {
	userDB := user.New(p.db.GetDB())
	salt, pass := hashGenerate.HashPassword(req.Password)
	userEntity := user.UserEntity{
		Login:    req.Name,
		Password: pass,
		Salt:     salt,
		Role:     user.UserRole(req.Role),
	}
	if err = userDB.Create(&userEntity); err != nil {
		res.Success = false
		res.Message = err.Error()
		return
	}
	res.Success = true
	res.Message = fmt.Sprintf("user %s created ID: %d", userEntity.Login, userEntity.ID)
	return
}

func (p userHandler) SignIn(ctx context.Context, logReq *pb.LoginRequest) (*pb.Token, error) {
	userDB := user.New(p.db.GetDB())
	user, err := userDB.FindByLogin(logReq.Name)
	if err != nil {
		return nil, err
	}
	if !hashGenerate.CheckPasswordHash(logReq.Password, user.Salt, user.Password) {
		return nil, fmt.Errorf("password incorect")
	}
	acToken, err := p.token.CreateToken(user.Role.String(), user.ID.String())
	if err != nil {
		return nil, err
	}
	return &pb.Token{AccessToken: acToken}, nil
}

func (p userHandler) CheckToken(ctx context.Context, token *pb.Token) (*pb.TokenResp, error) {
	claims, err := p.token.ParseTokenString(token.AccessToken)
	if err != nil {
		p.log.Warnln("ParseTokenString err:", err)
		return nil, err
	}
	return &pb.TokenResp{
		UserID: claims.UserID,
		Role:   pb.Role(pb.Role_value[claims.Role]),
	}, nil
}
