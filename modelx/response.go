package modelx

import (
	"errors"
	"net/http"
)

type Response struct {
	HTTPStatusCode int               `json:"-"`
	Code           string            `json:"code"`
	Description    string            `json:"description"`
	Reference      *Reference        `json:"reference,omitempty"`
	Data           interface{}       `json:"data,omitempty"`
	Header         map[string]string `json:"-"`
}

func NewResponse(code string) *Response {
	return &Response{
		Code:        code,
		Description: Description[code],
	}
}

func NewResponseError(httpStatusCode int, code string) *Response {
	return &Response{
		HTTPStatusCode: httpStatusCode,
		Code:           code,
		Description:    Description[code],
	}
}

func NewResponseSuccess(httpStatusCode int, data interface{}) *Response {
	return &Response{
		HTTPStatusCode: httpStatusCode,
		Code:           CodeSuccess,
		Description:    Description[CodeSuccess],
		Data:           data,
	}
}

func NewReponseBadRequest(desc string) (*Response, error) {
	return &Response{
			HTTPStatusCode: http.StatusBadRequest,
			Code:           CodeBadRequest,
			Description:    desc},
		errors.New(desc)
}
