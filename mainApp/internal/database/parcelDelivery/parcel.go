package parcelDelivery

import (
	"time"

	"gorm.io/gorm"
)

type ParcelDeliveryInter interface {
	Create(aPE *ParcelDeliveryEntity) error
	UpdateDestination(id int64, c Coordinates) error
}

func (a *ParcelDeliveryEntity) BeforeCreate(_ *gorm.DB) (err error) {
	a.Created = time.Now().Unix()
	return
}

type parcelDelivery struct {
	db *gorm.DB
}

func New(dbr *gorm.DB) ParcelDeliveryInter {

	return &parcelDelivery{db: dbr}
}

func (aP parcelDelivery) Create(pde *ParcelDeliveryEntity) error {
	if err := aP.db.Create(pde).Error; err != nil {
		return err
	}
	return nil
}

func (pD parcelDelivery) UpdateDestination(id int64, c Coordinates) error {
	if err := pD.db.Where("id", id).UpdateColumn("recipient_coordinate", c).Error; err != nil {
		return err
	}
	return nil
}
