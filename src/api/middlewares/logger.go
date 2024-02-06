package middlewares

import (
	"bytes"
	"io"
	"strings"
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
	return func(c *gin.Context) {
		if strings.Contains(c.FullPath(), "swagger") {
			c.Next()
		} else {
			blw := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			start := time.Now() // start
			path := c.FullPath()
			raw := c.Request.URL.RawQuery

			bodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body.Close()
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			c.Writer = blw
			c.Next()

			param := gin.LogFormatterParams{}
			param.TimeStamp = time.Now() // stop
			param.Latency = param.TimeStamp.Sub(start)
			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
			param.BodySize = c.Writer.Size()

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
}
