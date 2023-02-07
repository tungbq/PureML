package models

import (
	"fmt"
	"net/http"
	"reflect"
)

type ResponseBody struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`

	rawData interface{}
}

// RawData returns the unformatted error data (could be an internal error, text, etc.)
func (e *ResponseBody) RawData() any {
	return e.rawData
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
			Message: fmt.Sprintf("Internal Server Error - %s", err.Error()),
			Data:    nil,

			rawData: err,
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
	if data != nil && !(reflect.ValueOf(data).Kind() == reflect.Ptr && reflect.ValueOf(data).IsNil()) && reflect.TypeOf(data).Kind() != reflect.Slice {
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
