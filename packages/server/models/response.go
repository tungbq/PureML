package models

import "net/http"

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
