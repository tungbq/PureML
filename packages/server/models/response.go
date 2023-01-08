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

func NewErrorResponse(err error) *Response {
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
