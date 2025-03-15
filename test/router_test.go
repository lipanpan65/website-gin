package test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelloRoute(t *testing.T) {
	r := SetupTestRouter()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	w := PerformRequest(r, "GET", "/hello")

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 可以进一步检查响应体内容
	// 例如：
	// expected := `{"message": "Hello, World!"}`
	// if w.Body.String() != expected {
	//     t.Errorf("handler returned unexpected body: got %v want %v",
	//         w.Body.String(), expected)
	// }
}
