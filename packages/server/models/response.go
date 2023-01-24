package models

import (
	"net/http"
	"reflect"
)

type ResponseBody struct {
	Status  int
	Data    interface{}
	Message string
}

type Response struct {
	Error      error
	Body       ResponseBody
	StatusCode int
}
func (r *Response) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"status":  r.Body.Status,
		"data":    r.Body.Data,
		"message": r.Body.Message,
	}
}

func NewServerErrorResponse(err error) *Response {
	return &Response{
		Error: err,
		Body: ResponseBody{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		},
		StatusCode: http.StatusInternalServerError,
	}
}

func NewErrorResponse(statusCode int, message string) *Response {
	return &Response{
		Error: nil,
		Body: ResponseBody{
			Status:  statusCode,
			Message: message,
			Data:    nil,
		},
		StatusCode: statusCode,
	}
}

func NewDataResponse(statusCode int, data interface{}, message string) *Response {
	if reflect.TypeOf(data).Kind() != reflect.Slice {
		data = []interface{}{data}
	}
	return &Response{
		Error: nil,
		Body: ResponseBody{
			Status:  statusCode,
			Message: message,
			Data:    data,
		},
		StatusCode: statusCode,
	}
}
