package recovery

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGinRecoveryWithBrokenPipe(t *testing.T) {
	router := gin.New()
	router.Use(GinRecovery(true))
	router.GET("/testpath", func(c *gin.Context) {
		err := &net.OpError{
			Err: &os.SyscallError{
				Err:     errors.New("broken pipe"),
				Syscall: "write",
			},
		}
		panic(err)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/testpath", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGinRecoveryWithNonBrokenPipe(t *testing.T) {
	router := gin.Default()
	router.Use(GinRecovery(true))
	router.GET("/testpath", func(c *gin.Context) {
		panic(errors.New("non broken pipe error"))
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/testpath", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGinRecoveryWithoutPanic(t *testing.T) {
	router := gin.Default()
	router.Use(GinRecovery(true))
	router.GET("/testpath", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Test")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/testpath", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello Test", w.Body.String())
}
