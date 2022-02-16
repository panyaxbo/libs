package errorx

import (
	"errors"
	"net/http"
)

type HttpResponseError struct {
	StatusCode int

	Err error
}

func (r *HttpResponseError) Error() string {
	return r.Err.Error()
}

func (r *HttpResponseError) Temporary() bool {
	return r.StatusCode == http.StatusServiceUnavailable // 503
}

func do200OKResponse() error {
	return &HttpResponseError{
		StatusCode: 200,
		Err:        errors.New("200"),
	}
}
func do201CreatedResponse() error {
	return &HttpResponseError{
		StatusCode: 201,
		Err:        errors.New("Created"),
	}
}
func do400BadRequestResponse() error {
	return &HttpResponseError{
		StatusCode: 400,
		Err:        errors.New("Bad Request"),
	}
}

func do401UnauthorizedResponse() error {
	return &HttpResponseError{
		StatusCode: 401,
		Err:        errors.New("Unauthorized"),
	}
}
func do403ForbiddenResponse() error {
	return &HttpResponseError{
		StatusCode: 403,
		Err:        errors.New("Forbidden"),
	}
}

func do404NotFoundResponse() error {
	return &HttpResponseError{
		StatusCode: 404,
		Err:        errors.New("NotFound"),
	}
}
func do405MethodNotAllowedResponse() error {
	return &HttpResponseError{
		StatusCode: 405,
		Err:        errors.New("Method Not Allowed"),
	}
}
func do500InternalServerErrorResponse() error {
	return &HttpResponseError{
		StatusCode: 500,
		Err:        errors.New("Internal Server Error"),
	}
}
func do501NotImplementedResponse() error {
	return &HttpResponseError{
		StatusCode: 501,
		Err:        errors.New("Not Implemented"),
	}
}
func do502BadGatewayResponse() error {
	return &HttpResponseError{
		StatusCode: 502,
		Err:        errors.New("Bad Gateway"),
	}
}
func do503ServiceUnavailableResponse() error {
	return &HttpResponseError{
		StatusCode: 503,
		Err:        errors.New("Service Unavailable"),
	}
}
