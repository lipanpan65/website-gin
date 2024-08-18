package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的路由
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(200, "值：%v", "您好Gin")
	})
	r.GET("/news", func(context *gin.Context) {
		context.String(http.StatusOK, "This is new page")
	})
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Server error...")
	}
}
