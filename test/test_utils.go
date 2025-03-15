package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// SetupTestRouter 初始化一个用于测试的 Gin 引擎
func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

// PerformRequest 执行一个 HTTP 请求并返回响应记录器
func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
