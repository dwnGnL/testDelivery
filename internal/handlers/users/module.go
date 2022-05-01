package users

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mainApp/internal/models"
	"mainApp/pkg/config"
	"mainApp/pkg/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
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
}

func NewUserHandler(params Params) UserHandler {
	return &userHandler{db: params.DbInter, log: params.Logger, conf: params.Tuner}
}

func (p userHandler) SignUp(c *gin.Context) {
	var userReq models.UserCreateReq
	if err := c.ShouldBindJSON(&userReq); err != nil {
		p.log.Warnln("bind error")
		c.JSON(http.StatusBadGateway, gin.H{"error": "bind error"})
		return
	}

	// repo := actionPlan.New(p.db.GetDB())
	// if err := repo.Create(&acEntity); err != nil {
	// 	p.log.Warnln("can't create action plan entity", err)
	// 	c.JSON(http.StatusBadGateway, gin.H{"error": "can't create action plan"})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"created_id": ""})
}

func (p userHandler) sendRequest(method, uri string, reader io.Reader, respStruct interface{}, headers *map[string]string) error {
	req, err := http.NewRequest(method, p.conf.Task.Addr+uri, reader)
	if err != nil {
		return err
	}
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	// req.Header.Set("Content-Type")
	if headers != nil {
		for s, v := range *headers {
			req.Header.Set(s, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		if respStruct != nil {
			err = json.Unmarshal(body, &respStruct)

			if err != nil {
				return err
			}
		}

	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(body))
	}

	return nil
}
