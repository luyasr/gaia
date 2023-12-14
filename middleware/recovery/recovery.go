package recovery

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/log/zerolog"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

var (
	lastError      string
	lastErrorLower string
)

func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				brokenPipe := isBrokenPipe(err)

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zerolog.ConsoleLogger.Error().
						Any("errors", err).
						Str("request", string(httpRequest)).
						Send()
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: err check
					c.Abort()
					return
				}

				if stack {
					zerolog.ConsoleLogger.Error().
						Any("errors", err).
						Str("request", string(httpRequest)).
						Str("stack", string(debug.Stack())).
						Send()
				} else {
					zerolog.ConsoleLogger.Error().
						Any("errors", err).
						Str("request", string(httpRequest)).
						Send()
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
