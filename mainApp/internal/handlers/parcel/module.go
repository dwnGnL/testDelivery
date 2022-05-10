package parcel

import (
	"net/http"
	"testDelivery/mainApp/internal/models"
	"testDelivery/mainApp/pkg/config"
	"testDelivery/mainApp/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Provide(NewParcelHandler)

type ParcelHandler interface {
	Create(c *gin.Context, userID string)
}

type Params struct {
	fx.In
	db.DbInter
	*config.Tuner
	*logrus.Logger
}

type parcelHandler struct {
	db   db.DbInter
	log  *logrus.Logger
	conf *config.Tuner
}

func NewParcelHandler(params Params) ParcelHandler {

	return &parcelHandler{log: params.Logger, conf: params.Tuner, db: params.DbInter}
}

func (p parcelHandler) Create(c *gin.Context, userID string) {
	var parcelReq models.ParcelCreateReq
	if err := c.ShouldBindJSON(&parcelReq); err != nil {
		p.log.Warnln("bind error")
		c.JSON(http.StatusBadGateway, gin.H{"error": "bind error"})
		return
	}

}
