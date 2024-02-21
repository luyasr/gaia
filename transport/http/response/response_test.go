package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/luyasr/gaia/errors"
)

func TestGinJson(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GinJson(c, "test data")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"code":200`)
	assert.Contains(t, w.Body.String(), `"message":"success"`)
	assert.Contains(t, w.Body.String(), `"data":"test data"`)
}

func TestGinJsonWithError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GinJsonWithError(c, errors.New(10000, "test error", "test reason"))

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), `"code":400`)
	assert.Contains(t, w.Body.String(), `"message":"test error"`)
}
