package token

import (
	"fmt"
	"testDelivery/authorizationService/pkg/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(Setup)

type Params struct {
	fx.In
	*logrus.Logger
	*config.Tuner
}

type TokenInter interface {
	ParseTokenString(tokenString string) (JwtClaims, error)
	CreateToken(role, userID string) (string, error)
}

type jwtEntity struct {
	log        *logrus.Logger
	key        string
	ExpiredSec int64
}

func Setup(params Params) TokenInter {
	return &jwtEntity{
		log:        params.Logger,
		key:        params.Token.Key,
		ExpiredSec: params.Token.ExpiredSec,
	}
}

type JwtClaims struct {
	ExpiredAt int64  `json:"expired_at"`
	Role      string `json:"role"`
	UserID    string `json:"user_id"`
}

func (j JwtClaims) Valid() error {
	if j.ExpiredAt < time.Now().Unix() {
		return fmt.Errorf("token is expired")
	}
	return nil
}

func (t jwtEntity) CreateToken(role, userID string) (string, error) {
	var customFields JwtClaims
	customFields.ExpiredAt = time.Now().Unix()
	customFields.Role = role
	customFields.UserID = userID
	atToken := jwt.NewWithClaims(jwt.SigningMethodHS256, customFields)

	ss, err := atToken.SignedString(t.key)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (t jwtEntity) ParseTokenString(tokenString string) (JwtClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.key), nil
	})
	if err != nil {
		return JwtClaims{}, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !token.Valid || !ok {
		return JwtClaims{}, fmt.Errorf("token is invalid")
	}

	return *claims, nil
}
