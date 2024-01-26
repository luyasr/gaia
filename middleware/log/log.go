package log

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/log/zerolog"
)

type Log struct {
	logger *log.Helper
}

type Option func(*Log)

func Logger(logger *log.Helper) Option {
	return func(l *Log) {
		l.logger = logger
	}
}

func New(opts ...Option) *Log {
	l := &Log{
		logger: log.NewHelper(zerolog.New(zerolog.DefaultLogger)),
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

func (l *Log) GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start).String()

		logParams := l.getLogParams(c, cost)

		if len(c.Errors) > 0 {
			logParams = append(logParams, "errors", c.Errors.ByType(gin.ErrorTypePrivate).String())
		}

		l.logger.Infow(logParams...)
	}
}

func (l *Log) getLogParams(c *gin.Context, cost string) []any {
	return []any{
		"status", c.Writer.Status(),
		"method", c.Request.Method,
		"path", c.Request.URL.Path,
		"query", c.Request.URL.RawQuery,
		"ip", c.ClientIP(),
		"cost", cost,
	}
}
