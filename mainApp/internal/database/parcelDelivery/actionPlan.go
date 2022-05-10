package parcelDelivery

import (
	"time"

	"gorm.io/gorm"
)

type ParcelDeliveryInter interface {
	Create(aPE *ParcelDeliveryEntity) error
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
