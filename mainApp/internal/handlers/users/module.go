package users

import (
	"fmt"
	"log"
	"net/http"
	pb "testDelivery/authorizationProto"
	"testDelivery/mainApp/internal/models"
	"testDelivery/mainApp/pkg/config"
	"testDelivery/mainApp/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Provide(NewUserHandler)

type UserHandler interface {
	SignUp(c *gin.Context)
}

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
	pb.AuthorithationClient
}

func NewUserHandler(params Params) UserHandler {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", params.Tuner.Authorization.Addr, params.Tuner.Authorization.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to authoriization server: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuthorithationClient(conn)

	return &userHandler{db: params.DbInter, log: params.Logger, conf: params.Tuner, AuthorithationClient: c}
}

func (p userHandler) SignUp(c *gin.Context) {
	var userReq models.UserCreateReq
	if err := c.ShouldBindJSON(&userReq); err != nil {
		p.log.Warnln("bind error")
		c.JSON(http.StatusBadGateway, gin.H{"error": "bind error"})
		return
	}
	if _, ok := pb.Role_value[userReq.Role]; !ok {
		p.log.Warnln("role find err")
		c.JSON(http.StatusBadGateway, gin.H{"error": "not found this role"})
		return
	}
	authUserReq := pb.UserRequest{
		Name:     userReq.Username,
		Password: userReq.Password,
		Role:     pb.Role(pb.Role_value[userReq.Role]),
	}

	res, err := p.AuthorithationClient.SignUp(c.Request.Context(), &authUserReq)
	if err != nil || res == nil {
		p.log.Warnln("SignUp err")
		c.JSON(http.StatusBadGateway, gin.H{"error": "SignUp err " + err.Error()})
		return
	}

	if res.Success {
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": res.Message})
}

func (p userHandler) SignIn(c *gin.Context) {
	var userReq models.UserCreateReq
	if err := c.ShouldBindJSON(&userReq); err != nil {
		p.log.Warnln("bind error")
		c.JSON(http.StatusBadGateway, gin.H{"error": "bind error"})
		return
	}

	authUserReq := pb.LoginRequest{
		Name:     userReq.Username,
		Password: userReq.Password,
	}

	res, err := p.AuthorithationClient.SignIn(c.Request.Context(), &authUserReq)
	if err != nil || res == nil {
		p.log.Warnln("SignUp err")
		c.JSON(http.StatusBadGateway, gin.H{"error": "SignUp err " + err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"access_token": res.AccessToken})
}
