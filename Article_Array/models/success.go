package models

type ResponseSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
