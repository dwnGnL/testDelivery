package parcelDelivery

import (
	"time"

	"gorm.io/gorm"
)

type ParcelDeliveryInter interface {
	Create(aPE *ParcelDeliveryEntity) error
	UpdateDestination(id int64, c Coordinates) error
	CancelParcel(id int64) error
	Get(id int64) (ParcelDeliveryEntity, error)
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
	if err := pD.db.Model(ParcelDeliveryEntity{}).Where("id", id).UpdateColumn("recipient_coordinate", c).Error; err != nil {
		return err
	}
	return nil
}

func (pD parcelDelivery) CancelParcel(id int64) error {
	if err := pD.db.Where("id", id).UpdateColumn("status", IsCanceled).Error; err != nil {
		return err
	}
	return nil
}
func (aP parcelDelivery) Get(id int64) (ParcelDeliveryEntity, error) {
	var pde ParcelDeliveryEntity
	if err := aP.db.Find(&pde, id).Error; err != nil {
		return ParcelDeliveryEntity{}, err
	}
	return pde, nil
}
