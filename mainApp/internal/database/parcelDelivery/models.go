package parcelDelivery

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type ParcelDeliveryEntity struct {
	ID                  int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name                string         `gorm:"column:name"`
	Description         string         `gorm:"column:description"`
	RecipientCoordinate coordinates    `gorm:"column:recipient_coordinate;type:jsonb"`
	SenderCoordinate    coordinates    `gorm:"column:sender_coordinate;type:jsonb"`
	AdditionalInfo      AdditionalInfo `gorm:"column:additional_info;type:jsonb"`
	Created             int64          `gorm:"column:created"`
}

type AdditionalInfo map[string]interface{}

func (a AdditionalInfo) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *AdditionalInfo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type coordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (a coordinates) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *coordinates) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func (ParcelDeliveryEntity) TableName() string {
	return "parcel_delivery"
}
