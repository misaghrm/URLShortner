package model

type ResponseModel struct {
	Data    interface{} `json:"data,omitempty"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
}
