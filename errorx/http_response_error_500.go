package errorx

import (
	"errors"
	"fmt"
	"net/http"
)

type HttpResponse500Error struct {
	StatusCode int

	Err error
}

func (r *HttpResponse500Error) Error() string {
	return r.Err.Error()
}

func (r *HttpResponse500Error) Temporary() bool {
	return r.StatusCode == http.StatusServiceUnavailable
}

func (r *HttpResponse500Error) InternalServerError500Response() error {
	msg := fmt.Sprintf("%d - %s", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	return &HttpResponse500Error{
		StatusCode: http.StatusInternalServerError,
		Err:        errors.New(msg),
	}
}
