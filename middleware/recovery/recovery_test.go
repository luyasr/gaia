package recovery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/log/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestGinRecovery(t *testing.T) {
	// 创建一个新的Gin引擎
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// 创建一个新的日志助手
	logger := log.NewHelper(zerolog.New(zerolog.DefaultLogger))

	// 创建一个新的Recovery实例
	r := New(Logger(logger))

	// 添加GinRecovery中间件
	router.Use(r.GinRecovery(true))

	// 添加一个测试路由
	router.GET("/test", func(c *gin.Context) {
		panic("Test panic")
	})

	// 创建一个新的HTTP请求
	req, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// 创建一个HTTP响应记录器
	w := httptest.NewRecorder()

	// 执行HTTP请求
	router.ServeHTTP(w, req)

	// 检查响应状态码
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}