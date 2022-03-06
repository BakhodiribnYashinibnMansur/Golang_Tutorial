package models

type DefaultError struct {
	Message string `json:"message"`
}

type ResponseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
