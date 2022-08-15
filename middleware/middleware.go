package middleware

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/panyaxbo/libs/configx"
	"github.com/panyaxbo/libs/contextx"
	"github.com/panyaxbo/libs/logx"
	"github.com/panyaxbo/libs/maskx"
	"github.com/sirupsen/logrus"
)

var (
	// DefaultSkipper default of skipper
	DefaultSkipper = func(c echo.Context) bool {
		return c.Path() == "/health"
	}
)

// Skipper skip middleware
type Skipper func(c echo.Context) bool

// Recover returns a middleware which recovers from panics anywhere in the chain
// and handles the control to the centralized HTTPErrorHandler.
func Recover() echo.MiddlewareFunc {
	return RecoverWithConfig(middleware.DefaultRecoverConfig)
}

func RecoverWithConfig(config middleware.RecoverConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultRecoverConfig.Skipper
	}
	if config.StackSize == 0 {
		config.StackSize = middleware.DefaultRecoverConfig.StackSize
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					stack := make([]byte, config.StackSize)
					length := runtime.Stack(stack, !config.DisableStackAll)
					if !config.DisablePrintStack {
						logx.WithContext(c.Request().Context()).Errorf("[PANIC RECOVER] %v %s\n", err, stack[:length])
					}
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}

// RequestID returns a X-Request-ID middleware.
func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if DefaultSkipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			rid := req.Header.Get(echo.HeaderXRequestID)
			if rid == "" {
				rid = uuid.New().String()
			}
			res.Header().Set(echo.HeaderXRequestID, rid)

			ctx := contextx.SetID(req.Context(), rid)
			c.SetRequest(req.WithContext(ctx))

			return next(c)
		}
	}
}

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if DefaultSkipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			ctx := req.Context()

			b := make([]byte, 0)
			if req.Body != nil {
				b, _ = ioutil.ReadAll(req.Body)
			}
			req.Body = ioutil.NopCloser(bytes.NewBuffer(b))

			logx.WithContext(ctx).WithFields(logrus.Fields{
				"header": req.Header,
				"body":   logx.LimitMSGByte(b),
			}).Info("echo request information")

			resBody := new(bytes.Buffer)
			mw := io.MultiWriter(res.Writer, resBody)
			writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: res.Writer}
			res.Writer = writer

			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}

			duration := time.Since(start)

			logx.WithContext(ctx).WithFields(logrus.Fields{
				"header":          res.Header(),
				"body":            logx.LimitMSGByte(resBody.Bytes()),
				"method":          req.Method,
				"host":            req.Host,
				"path_uri":        req.RequestURI,
				"remote_ip":       c.RealIP(),
				"status":          res.Status,
				"duration_string": duration.String(),
				"duration":        duration,
			}).Info("echo response information")

			return nil
		}
	}
}
func LoggerWithMaskLog(config *configx.ConfigMaskLog) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if DefaultSkipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			ctx := req.Context()

			b := make([]byte, 0)
			if req.Body != nil {
				b, _ = ioutil.ReadAll(req.Body)
			}
			req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			//var maskingInstance = NewMaskTool()
			//	maskTool := NewMaskTool(filter.FieldFilter("identifier"))

			if config.IsMaskLogWithEncrypt {
				if req.Body != nil || string(b) != "" {
					m := maskx.Init(configx.SensitiveFields)
					t, err := m.JsonMaskEncrypted(b, config.Env)
					if err != nil {
						logx.WithContext(ctx).Panicf("%s", err)
					}
					logx.WithContext(ctx).WithFields(logrus.Fields{
						"header": req.Header,
						"body":   logx.LimitMSGByte([]byte(*t)),
					}).Info("echo request information")
				} else {
					logx.WithContext(ctx).WithFields(logrus.Fields{
						"header": req.Header,
						"body":   logx.LimitMSGByte(b),
					}).Info("echo request information")
				}
			} else if config.IsMaskLogWithSymbol {
				if req.Body != nil || string(b) != "" {
					m := maskx.Init(configx.SensitiveFields)
					t, err := m.JsonMaskSymbol(b, config.Symbol)
					if err != nil {
						logx.WithContext(ctx).Panicf("%s", err)
					}
					logx.WithContext(ctx).WithFields(logrus.Fields{
						"header": req.Header,
						"body":   logx.LimitMSGByte([]byte(*t)),
					}).Info("echo request information")
				} else {
					logx.WithContext(ctx).WithFields(logrus.Fields{
						"header": req.Header,
						"body":   logx.LimitMSGByte(b),
					}).Info("echo request information")
				}

			} else {
				logx.WithContext(ctx).WithFields(logrus.Fields{
					"header": req.Header,
					"body":   logx.LimitMSGByte(b),
				}).Info("echo request information")
			}
			resBody := new(bytes.Buffer)
			mw := io.MultiWriter(res.Writer, resBody)
			writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: res.Writer}
			res.Writer = writer

			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}

			duration := time.Since(start)

			if config.IsMaskLogWithEncrypt {
				m := maskx.Init(sensitiveFields)
				t, err := m.JsonMaskEncrypted(resBody.Bytes(), config.Env)
				if err != nil {
					logx.WithContext(ctx).Panicf("%s", err)
				}
				logx.WithContext(ctx).WithFields(logrus.Fields{
					"header":          res.Header(),
					"body":            logx.LimitMSGByte([]byte(*t)),
					"method":          req.Method,
					"host":            req.Host,
					"path_uri":        req.RequestURI,
					"remote_ip":       c.RealIP(),
					"status":          res.Status,
					"duration_string": duration.String(),
					"duration":        duration,
				}).Info("echo response information")
			} else if config.IsMaskLogWithSymbol {
				m := maskx.Init(sensitiveFields)
				t, err := m.JsonMaskSymbol(resBody.Bytes(), config.Symbol)
				if err != nil {
					logx.WithContext(ctx).Panicf("%s", err)
				}
				logx.WithContext(ctx).WithFields(logrus.Fields{
					"header":          res.Header(),
					"body":            logx.LimitMSGByte([]byte(*t)),
					"method":          req.Method,
					"host":            req.Host,
					"path_uri":        req.RequestURI,
					"remote_ip":       c.RealIP(),
					"status":          res.Status,
					"duration_string": duration.String(),
					"duration":        duration,
				}).Info("echo response information")
			} else {
				logx.WithContext(ctx).WithFields(logrus.Fields{
					"header":          res.Header(),
					"body":            logx.LimitMSGByte(resBody.Bytes()),
					"method":          req.Method,
					"host":            req.Host,
					"path_uri":        req.RequestURI,
					"remote_ip":       c.RealIP(),
					"status":          res.Status,
					"duration_string": duration.String(),
					"duration":        duration,
				}).Info("echo response information")
			}
			return nil
		}
	}
}

func Health() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Path() == "/health" {
				return c.JSON(http.StatusOK, "ok!")
			}
			return next(c)
		}
	}
}

func JWT(i interface{}) echo.MiddlewareFunc {
	return middleware.JWT(i)
}

func CORS() echo.MiddlewareFunc {
	return middleware.CORS()
}

func Secure() echo.MiddlewareFunc {
	return middleware.Secure()
}

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyDumpResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *bodyDumpResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *bodyDumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
