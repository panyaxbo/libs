package errorx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

const (
	// BadRequestCode is error code when request is invalid
	BadRequestCode = "20000"

	// InternalServerErrorCode is error code when error occur in app
	InternalServerErrorCode = "80000"

	// ExternalErrorCode is error code when other api return unexpected result
	ExternalErrorCode = "80001"

	// ExternalTimeoutCode is error code for other api timeout
	ExternalTimeoutCode = "80002"

	// SuccessCode is code when api process successfully
	SuccessCode = "0"
)

func NewBadRequest(msg, description string, cause ...error) error {
	return New(http.StatusBadRequest, BadRequestCode, msg, description, cause...)
}

func NewInternalServerError(msg, description string, cause ...error) error {
	return New(http.StatusInternalServerError, InternalServerErrorCode, msg, description, cause...)
}

func NewExternalError(msg, description string, cause ...error) error {
	return New(http.StatusInternalServerError, ExternalErrorCode, msg, description, cause...)
}

func NewExternalTimeout(msg, description string, cause ...error) error {
	return New(http.StatusInternalServerError, ExternalTimeoutCode, msg, description, cause...)
}

// func HTTPErrorHandler(err error, c echo.Context) {
// 	if _, ok := errors.Cause(err).(*Errs); ok {
// 		JSON(c, err)
// 		return
// 	}

// 	code := http.StatusInternalServerError
// 	if he, ok := err.(*echo.HTTPError); ok {
// 		code = he.Code
// 	}

// 	JSON(c, New(code, InternalServerErrorCode, "general", err.Error(), err))
// }

// func JSON(c echo.Context, err error) error {
// 	if errs, ok := errors.Cause(err).(*Errs); ok {
// 		lang := c.Request().Header.Get("Accept-Language")
// 		MapErrorDetail(errs, lang)
// 		logx.ErrorfWithID(c, "[RAW ERROR] %+v", err)
// 		return c.JSON(errs.HTTPStatusCode, errs)
// 	}

// 	logx.ErrorfWithID(
// 		c,
// 		fmt.Sprint(c.Path(), " [RAW ERROR] %+v"),
// 		err,
// 	)
// 	return c.JSON(http.StatusInternalServerError, errors.Cause(NewInternalServerError("general", err.Error(), err)))
// }

// func MapErrorDetail(errs *Errs, lang string) {
// 	if consul.ConfigMap != nil {
// 		errorHandler := newErrorHandler(consul.ConfigMap)
// 		if errorStatusLookup, err := errorHandler.getError(errs.Code); err == nil {
// 			if lang == "th-TH" {
// 				errs.Title = overrideMsg(errorStatusLookup.TitleTh, errs.Title, errs.isOverride)
// 				errs.Msg = overrideMsg(errorStatusLookup.MessageTh, errs.Msg, errs.isOverride)
// 				errs.Description = overrideMsg(errorStatusLookup.Description, errs.Description, errs.isOverride)
// 			} else {
// 				errs.Title = overrideMsg(errorStatusLookup.TitleEn, errs.Title, errs.isOverride)
// 				errs.Msg = overrideMsg(errorStatusLookup.MessageEn, errs.Msg, errs.isOverride)
// 				errs.Description = overrideMsg(errorStatusLookup.Description, errs.Description, errs.isOverride)
// 			}
// 		}
// 	}
// }

func New(httpStatusCode int, externalErrorCode, msg, description string, cause ...error) error {
	var inner error
	var isOverride bool
	if len(cause) > 0 {
		if errsOverride, ok := cause[0].(*errsOverride); ok {
			isOverride = errsOverride.isOverride
			inner = errsOverride.cause
		} else {
			inner = cause[0]
		}
	}

	err := &Errs{
		HTTPStatusCode: httpStatusCode,
		Code:           externalErrorCode,
		Msg:            msg,
		Description:    description,
		Cause:          inner,
		isOverride:     isOverride,
	}

	if _, ok := inner.(stackTracer); ok {
		return err
	}

	return errors.Wrap(err, msg)
}

func NewFail(httpStatusCode int, externalErrorCode, msg, description, failAttempt, maxFail string, cause ...error) error {
	var inner error
	var isOverride bool
	if len(cause) > 0 {
		if errsOverride, ok := cause[0].(*errsOverride); ok {
			isOverride = errsOverride.isOverride
			inner = errsOverride.cause
		} else {
			inner = cause[0]
		}
	}

	err := &Errs{
		HTTPStatusCode: httpStatusCode,
		Code:           externalErrorCode,
		Msg:            fmt.Sprintf("%s (%s/%s)", msg, failAttempt, maxFail),
		Description:    description,
		Cause:          inner,
		isOverride:     isOverride,
	}

	if _, ok := inner.(stackTracer); ok {
		return err
	}

	return errors.Wrap(err, msg)
}

func NewWithAdditionalInfo(httpStatusCode int, internalErrorCode, msg, description string, additionalInfo interface{}, cause ...error) error {
	var inner error
	var isOverride bool
	if len(cause) > 0 {
		if errsOverride, ok := cause[0].(*errsOverride); ok {
			isOverride = errsOverride.isOverride
			inner = errsOverride.cause
		} else {
			inner = cause[0]
		}
	}

	err := &Errs{
		HTTPStatusCode: httpStatusCode,
		Code:           internalErrorCode,
		Msg:            msg,
		Description:    description,
		Cause:          inner,
		AdditionalInfo: &additionalInfo,
		isOverride:     isOverride,
	}

	if _, ok := inner.(stackTracer); ok {
		return err
	}

	return errors.Wrap(err, msg)
}

//  * External Error Code is error code from legacy services such as CBPAY, CBS etc., For instance 0000, 5431 etc.
//  * Internal Error Code is our error code mapping, For instance 80000, 82000 etc.
// func NewWithExternalErrorCode(httpStatusCode int, errorGroup ErrorGroup, externalErrorCode, msg, description string, additionalInfo interface{}, cause ...error) error {
// 	var inner error
// 	var isOverride bool
// 	if len(cause) > 0 {
// 		if errsOverride, ok := cause[0].(*errsOverride); ok {
// 			isOverride = errsOverride.isOverride
// 			inner = errsOverride.cause
// 		} else {
// 			inner = cause[0]
// 		}
// 	}

// 	internalErrorCode := GenericErrorCode
// 	var err error
// 	if consul.ConfigMap != nil {
// 		errorHandler := newErrorHandler(consul.ConfigMap)
// 		internalErrorCode, err = errorHandler.getInternalErrorCode(errorGroup, externalErrorCode)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	err = &Errs{
// 		HTTPStatusCode: httpStatusCode,
// 		Code:           internalErrorCode,
// 		Msg:            msg,
// 		Description:    description,
// 		Cause:          inner,
// 		AdditionalInfo: &additionalInfo,
// 		isOverride:     isOverride,
// 	}

// 	if _, ok := inner.(stackTracer); ok {
// 		return err
// 	}

// 	return errors.Wrap(err, msg)
// }

// Support QR Xborder only.
// func NewWithReasonCode(httpStatusCode int, reasonCode, msg, description string, additionalInfo interface{}, cause ...error) error {
// 	var inner error
// 	var isOverride bool
// 	if len(cause) > 0 {
// 		if errsOverride, ok := cause[0].(*errsOverride); ok {
// 			isOverride = errsOverride.isOverride
// 			inner = errsOverride.cause
// 		} else {
// 			inner = cause[0]
// 		}
// 	}

// 	internalErrorCode := GenericErrorCode
// 	var err error
// 	if consul.ConfigMap != nil {
// 		errorHandler := newErrorHandler(consul.ConfigMap)
// 		internalErrorCode, err = errorHandler.getInternalErrorCodeWithReasonCode(reasonCode)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	err = &Errs{
// 		HTTPStatusCode: httpStatusCode,
// 		Code:           internalErrorCode,
// 		Msg:            msg,
// 		Description:    description,
// 		Cause:          inner,
// 		AdditionalInfo: &additionalInfo,
// 		isOverride:     isOverride,
// 	}

// 	if _, ok := inner.(stackTracer); ok {
// 		return err
// 	}

// 	return errors.Wrap(err, msg)
// }

// func GetInternalErrorCode(errorGroup ErrorGroup, externalErrorCode string) (string, error) {
// 	errorHandler := newErrorHandler(consul.ConfigMap)
// 	return errorHandler.getInternalErrorCode(errorGroup, externalErrorCode)
// }

// // actual money out group - validate error/timeout with status code 80008(error)/80009(error)
// func ValidateMoneyOut(err error) error {
// 	if err, ok := err.(net.Error); ok && err.Timeout() {
// 		return New(http.StatusRequestTimeout, MoneyOutTimeoutCode, "", "")
// 	}

// 	return New(http.StatusInternalServerError, MoneyOutErrorCode, "", "")
// }

// func ValidateTimeOut(err error) error {
// 	if err, ok := err.(net.Error); ok && err.Timeout() {
// 		return New(http.StatusRequestTimeout, TimeOutErrorCode, "", "")
// 	}

// 	return err
// }

func NewErrorObject(httpStatusCode int, code, msg, description string, cause ...error) []byte {
	var inner error
	var isOverride bool
	if len(cause) > 0 {
		if errsOverride, ok := cause[0].(*errsOverride); ok {
			isOverride = errsOverride.isOverride
			inner = errsOverride.cause
		} else {
			inner = cause[0]
		}
	}
	detailError := &Errs{
		HTTPStatusCode: httpStatusCode,
		Code:           code,
		Msg:            msg,
		Description:    description,
		Cause:          inner,
		isOverride:     isOverride,
	}
	resdetail, _ := json.Marshal(detailError)
	return resdetail

}

type Errs struct {
	HTTPStatusCode int          `json:"-"`
	Code           string       `json:"code"`
	Msg            string       `json:"message"`
	Description    string       `json:"description"`
	Cause          error        `json:"-"`
	Title          string       `json:"title"`
	AdditionalInfo *interface{} `json:"additionalInfo"`
	isOverride     bool
}

type AdditionalInfo struct {
	HostCode        *string `json:"hostCode"`
	HostDescription *string `json:"hostDescription"`
	EntityInfo      *string `json:"entityInfo"`
}

type errsOverride struct {
	isOverride bool
	cause      error
}

func IsOverride(isOverride bool, cause ...error) *errsOverride {
	var inner error
	if len(cause) > 0 {
		inner = cause[0]
	}
	return &errsOverride{
		isOverride: isOverride,
		cause:      inner,
	}
}

func (e *Errs) Error() string {
	if e.Cause == nil {
		return fmt.Sprintf("code:%s, msg:%s, description:%s", e.Code, e.Msg, e.Description)
	}
	if st, ok := e.Cause.(stackTracer); ok {
		stack := strings.Replace(fmt.Sprintf("\t%+v", st), "\n", "\n\t", -1)
		return fmt.Sprintf("code:%s, msg:%s, description:%s, isOverride:%t\ncause:\n%+v", e.Code, e.Msg, e.Description, e.isOverride, stack)
	}
	return fmt.Sprintf("code:%s, msg:%s, description:%s, isOverride:%t, cause:\n%+v", e.Code, e.Msg, e.Description, e.isOverride, e.Cause)
}

func (e *errsOverride) Error() string {
	if e.cause != nil {
		if st, ok := e.cause.(stackTracer); ok {
			stack := strings.Replace(fmt.Sprintf("\t%+v", st), "\n", "\n\t", -1)
			return fmt.Sprintf("isOverride:%t\ncause:\n%+v", e.isOverride, stack)
		}
	}
	return fmt.Sprintf("isOverride:%t, cause:\n%+v", e.isOverride, e.cause)
}

func (e *Errs) StackTrace() errors.StackTrace {
	if st, ok := e.Cause.(stackTracer); ok {
		return st.StackTrace()
	}
	return nil
}

func overrideMsg(src, msg string, isOverride bool) string {
	if isOverride && msg != "" {
		return msg
	}
	return src
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}
