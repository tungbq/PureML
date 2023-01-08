package models

type Response struct {
	Error      error
	Message    string
	Body       interface{}
	StatusCode int
}

func NewErrorResponse(err error) *Response {
	return &Response{
		Error: err,
	}
}
