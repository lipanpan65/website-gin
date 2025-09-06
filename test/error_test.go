package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"website-gin/utils"
	"website-gin/utils/errors"
)

func TestResultError(t *testing.T) {
	r := SetupTestRouter()
	r.GET("/test-error", func(c *gin.Context) {
		utils.ResultError(c, errors.DictExisted)
	})

	w := PerformRequest(r, "GET", "/test-error")

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 可以进一步检查响应体中的错误信息
	// 例如：
	// expected := `{"success": false, "code": 400, "message": "字典已经存在", "data": null}`
	// if w.Body.String() != expected {
	//     t.Errorf("handler returned unexpected body: got %v want %v",
	//         w.Body.String(), expected)
	// }
}
