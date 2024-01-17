package middlewares

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/pkg/logging"
)

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *BodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func DefaultStracturedLogger(cfg *config.Config) gin.HandlerFunc {
	logger := logging.NewLogger(cfg)
	return stractureLogger(logger)
}

func stractureLogger(logger logging.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blw := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		start := time.Now()
		path := ctx.FullPath()
		raw := ctx.Request.URL.RawQuery

		bodyBytes, _ := ioutil.ReadAll(ctx.Request.Body)
		ctx.Request.Body.Close()
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		ctx.Writer = blw
		ctx.Next()

		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = ctx.ClientIP()
		param.Method = ctx.Request.Method
		param.StatusCode = ctx.Writer.Status()
		param.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = ctx.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		keys := map[logging.ExtraKey]interface{}{}

		keys[logging.Path] = param.Path
		keys[logging.ClientIp] = param.ClientIP
		keys[logging.Method] = param.Method
		keys[logging.Latency] = param.Latency
		keys[logging.StatusCode] = param.StatusCode
		keys[logging.ErrorMessage] = param.ErrorMessage
		keys[logging.BodySize] = param.BodySize
		keys[logging.RequestBody] = string(bodyBytes)
		keys[logging.ResponseBody] = blw.body.String()

		logger.Info(logging.RequestResponse, logging.Api, "", keys)
	}
}
