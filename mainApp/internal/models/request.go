package models

type UserCreateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"` //for easy test
}

type ParcelCreateReq struct {
	NameOfItem          string                 `json:"name_of_item"`
	Description         string                 `json:"description"`
	RecipientCoordinate Coordinates            `json:"recipient_coordinate"`
	SenderCoordinate    Coordinates            `json:"sender_coordinate"`
	AdditionalInfo      map[string]interface{} `json:"additional_info"`
}

type Coordinates struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
