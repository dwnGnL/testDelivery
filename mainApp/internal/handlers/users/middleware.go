package users

import (
	"errors"
	"net/http"
	"strings"
	pb "testDelivery/authorizationProto"

	"github.com/gin-gonic/gin"
)

var Roles pb.Role

func (p userHandler) CheckAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenStr, err := jwtFromHeader(c, "Authorization")
		if err != nil {
			unauthorized(c, http.StatusUnauthorized, "auth header empty")
			return
		}

		tokenResp, err := p.AuthorithationClient.CheckToken(c.Request.Context(), &pb.Token{AccessToken: tokenStr})
		if err != nil {
			unauthorized(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Set("role", tokenResp.Role)
		c.Set("userID", tokenResp.UserID)
		c.Next()
		return
	}
}

func jwtFromHeader(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)

	if authHeader == "" {
		return "", errors.New("auth header empty")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("invalid auth header")
	}

	return parts[1], nil
}

func unauthorized(c *gin.Context, code int, message string) {
	c.Abort()
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

type funcGin func(c *gin.Context, userID string)

func PrepareHandler(f funcGin, accessFor ...pb.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userRole pb.Role
		var userID string
		roleinterface, exist := c.Get("role")
		if exist {
			userRole = roleinterface.(pb.Role)
		}
		userIDinterface, exist := c.Get("userID")
		if exist {
			userID = userIDinterface.(string)
		}
		for _, role := range accessFor {
			if role == userRole {
				f(c, userID)
				return
			}
		}
		unauthorized(c, http.StatusForbidden, "доступ запрещен")
		return
	}
}
