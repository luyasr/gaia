package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGinLogger(t *testing.T) {
	router := gin.Default()
	router.Use(GinLogger())
	router.GET("/testpath", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Test")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/testpath", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello Test", w.Body.String())
}

func TestGinLoggerWithError(t *testing.T) {
	router := gin.New()
	router.Use(GinLogger())
	router.GET("/testpath", func(c *gin.Context) {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/testpath", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "Internal Server Error", w.Body.String())
}

func TestGinLoggerWithLongProcessingTime(t *testing.T) {
	router := gin.New()
	router.Use(GinLogger())
	router.GET("/testpath", func(c *gin.Context) {
		time.Sleep(2 * time.Second)
		c.String(http.StatusOK, "Hello Test")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/testpath", nil)
	start := time.Now()
	router.ServeHTTP(w, req)
	elapsed := time.Since(start)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello Test", w.Body.String())
	assert.True(t, elapsed >= 2*time.Second)
}
