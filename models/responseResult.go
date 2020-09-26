package models

type ResponseResult struct {
	Code int `json:"code"`
	Message string`json:"message"`
	Date interface{}`json:"date"`
}
