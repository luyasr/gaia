package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/log/zerolog"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start)

		logEvent := zerolog.ConsoleLogger.Info().
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("query", c.Request.URL.RawQuery).
			Str("ip", c.ClientIP()).
			Str("cost", cost.String())

		if len(c.Errors) > 0 {
			logEvent.Str("errors", c.Errors.ByType(gin.ErrorTypePrivate).String())
		}

		logEvent.Send()
	}
}
