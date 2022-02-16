package errorx

import (
	"errors"
	"fmt"
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
	return r.StatusCode == http.StatusServiceUnavailable
}

func (r *HttpResponseError) do200OKResponse() error {
	msg := fmt.Sprintf("%d - %s", http.StatusOK, http.StatusText(http.StatusOK))
	return &HttpResponseError{
		StatusCode: http.StatusOK,
		Err:        errors.New(msg),
	}
}
func (r *HttpResponseError) do201CreatedResponse() error {
	msg := fmt.Sprintf("%d - %s", http.StatusCreated, http.StatusText(http.StatusCreated))
	return &HttpResponseError{
		StatusCode: http.StatusCreated,
		Err:        errors.New(msg),
	}
}
func (r *HttpResponseError) do400BadRequestResponse() error {
	msg := fmt.Sprintf("%d - %s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	return &HttpResponseError{
		StatusCode: http.StatusBadRequest,
		Err:        errors.New(msg),
	}
}

func (r *HttpResponseError) do401UnauthorizedResponse() error {
	msg := fmt.Sprintf("%d - %s", http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	return &HttpResponseError{
		StatusCode: http.StatusUnauthorized,
		Err:        errors.New(msg),
	}
}
func (r *HttpResponseError) do403ForbiddenResponse() error {
	msg := fmt.Sprintf("%d - %s", http.StatusForbidden, http.StatusText(http.StatusForbidden))
	return &HttpResponseError{
		StatusCode: http.StatusForbidden,
		Err:        errors.New(msg),
	}
}

func (r *HttpResponseError) do404NotFoundResponse() error {
	msg := fmt.Sprintf("%d - %s", http.StatusNotFound, http.StatusText(http.StatusNotFound))
	return &HttpResponseError{
		StatusCode: http.StatusNotFound,
		Err:        errors.New(msg),
	}
}
func (r *HttpResponseError) do405MethodNotAllowedResponse() error {
	msg := fmt.Sprintf("%d - %s", http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
	return &HttpResponseError{
		StatusCode: http.StatusMethodNotAllowed,
		Err:        errors.New(msg),
	}
}
func (r *HttpResponseError) do500InternalServerErrorResponse() error {
	msg := fmt.Sprintf("%d - %s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	return &HttpResponseError{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(msg),
	}
}
func (r *HttpResponseError) do501NotImplementedResponse() error {
	msg := fmt.Sprintf("%d - %s", http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
	return &HttpResponseError{
		StatusCode: http.StatusNotImplemented,
		Err:        errors.New(msg),
	}
}
func (r *HttpResponseError) do502BadGatewayResponse() error {
	msg := fmt.Sprintf("%d - %s", http.StatusBadGateway, http.StatusText(http.StatusBadGateway))
	return &HttpResponseError{
		StatusCode: http.StatusBadGateway,
		Err:        errors.New(msg),
	}
}
func (r *HttpResponseError) do503ServiceUnavailableResponse() error {
	msg := fmt.Sprintf("%d - %s", http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable))
	return &HttpResponseError{
		StatusCode: http.StatusServiceUnavailable,
		Err:        errors.New(msg),
	}
}
