package models

type Config struct {
	Main  ConfMain
	DB    ConfDB
	Token ConfToken
}

// ConfMain - basic configuration
type ConfMain struct {
	Port string
	Name string
}

type ConfToken struct {
	Key        string
	ExpiredSec int64
}
type DefaultConfForService struct {
	Addr string
	Port string
	Name string
}
type ConfDB struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSlMode  string
}

type Response struct {
	Success   string       `json:"success"`
	ErrorResp *ErrorStruct `json:"errorResp,omitempty"`
	Response  *interface{} `json:"response,omitempty"`
}

type ErrorStruct struct {
	ErrorCode    int           `json:"error_code"`
	ErrorMessage string        `json:"error_msg"`
	ErrorData    []interface{} `json:"error_data"`
}
