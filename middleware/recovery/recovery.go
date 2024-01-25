package recovery

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/log/zerolog"
)

var (
	lastError      string
	lastErrorLower string
)

type Recovery struct {
	logger *log.Helper
}

type Option func(*Recovery)

func WithLogger(logger *log.Helper) Option {
	return func(r *Recovery) {
		r.logger = logger
	}
}

func New(opts ...Option) *Recovery {
	r := &Recovery{
		logger: log.NewHelper(zerolog.New(zerolog.DefaultLogger)),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *Recovery) GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				brokenPipe := isBrokenPipe(err)

				if brokenPipe {
					r.logger.Errorw(
						"request", string(httpRequest),
						"errors", err,
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: err check
					c.Abort()
					return
				}

				if stack {
					r.logger.Errorw(
						"request", string(httpRequest),
						"errors", err,
						"stack", string(debug.Stack()),
					)
				} else {
					r.logger.Errorw(
						"request", string(httpRequest),
						"errors", err,
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func isBrokenPipe(err any) bool {
	if ne, ok := err.(*net.OpError); ok {
		if se, ok := ne.Err.(*os.SyscallError); ok {
			errStr := se.Error()
			if errStr != lastError {
				lastError = errStr
				lastErrorLower = strings.ToLower(lastError)
			}
			if strings.Contains(lastError, "broken pipe") || strings.Contains(lastError, "connection reset by peer") ||
				strings.Contains(lastErrorLower, "broken pipe") || strings.Contains(lastErrorLower, "connection reset by peer") {
				return true
			}
		}
	}
	return false
}