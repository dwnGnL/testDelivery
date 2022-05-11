package parcel

import (
	"net/http"
	"strconv"
	"testDelivery/mainApp/internal/database/parcelDelivery"
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
	pD := parcelDelivery.New(p.db.GetDB())
	pDEntity := parcelDelivery.ParcelDeliveryEntity{
		Name:                parcelReq.NameOfItem,
		Description:         parcelReq.Description,
		RecipientCoordinate: parcelDelivery.Coordinates(parcelReq.RecipientCoordinate),
		SenderCoordinate:    parcelDelivery.Coordinates(parcelReq.SenderCoordinate),
		AdditionalInfo:      parcelReq.AdditionalInfo,
	}
	if err := pD.Create(&pDEntity); err != nil {
		p.log.Warnln("bind error")
		c.JSON(http.StatusBadGateway, gin.H{"error": "bind error"})
		return
	}
	c.JSON(http.StatusOK, pDEntity)
}

func (p parcelHandler) ChangeDestination(c *gin.Context, userID string) {
	var parcelReq models.ParcelCreateReq
	if err := c.ShouldBindJSON(&parcelReq); err != nil {
		p.log.Warnln("bind error")
		c.JSON(http.StatusBadGateway, gin.H{"error": "bind error"})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		p.log.Warnln("parse id error")
		c.JSON(http.StatusBadGateway, gin.H{"error": "bind error"})
		return
	}
	pD := parcelDelivery.New(p.db.GetDB())

	if err := pD.UpdateDestination(id, parcelDelivery.Coordinates(parcelReq.RecipientCoordinate)); err != nil {
		p.log.Warnln("bind error")
		c.JSON(http.StatusBadGateway, gin.H{"error": "bind error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
