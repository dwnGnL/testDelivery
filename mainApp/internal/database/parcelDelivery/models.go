package parcelDelivery

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Status string

const (
	IsCreated    Status = "is_created"
	IsCanceled   Status = "is_canceled"
	IsAccepted   Status = "is_accepted"
	IsDelivering Status = "is_delivering"
	IsDelivered  Status = "is_delivered"
)

func (s Status) possibleStatussesTransitions() []Status {
	switch s {
	case IsCreated:
		return []Status{}
	case IsCanceled:
		return []Status{IsCreated, IsAccepted, IsDelivering}
	case IsAccepted:
		return []Status{IsCreated}
	case IsDelivering:
		return []Status{IsAccepted}
	case IsDelivered:
		return []Status{IsDelivering}

	}
	return []Status{}
}

func (s Status) CheckTransitionPossible(incomingStatus Status) bool {
	for _, v := range s.possibleStatussesTransitions() {
		if v == incomingStatus {
			return true
		}
	}
	return false
}

// var
func (s *Status) Scan(value interface{}) error {
	*s = Status(value.(string))
	return nil
}

func (s Status) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *Status) String() string {
	return string(*s)
}

type ParcelDeliveryEntity struct {
	ID                  int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	Name                string         `gorm:"column:name"`
	Description         string         `gorm:"column:description"`
	RecipientCoordinate Coordinates    `gorm:"column:recipient_coordinate;type:jsonb"`
	SenderCoordinate    Coordinates    `gorm:"column:sender_coordinate;type:jsonb"`
	AdditionalInfo      AdditionalInfo `gorm:"column:additional_info;type:jsonb"`
	Status              Status         `gorm:"column:status;type:parcel_status"`
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

type Coordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (a Coordinates) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Coordinates) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

func (ParcelDeliveryEntity) TableName() string {
	return "parcel_delivery"
}
