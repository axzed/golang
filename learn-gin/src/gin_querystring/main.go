package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(ctx *gin.Context) {
		// 获取浏览器发请求携带的query string参数
		// name := ctx.Query("query")
		// DefaultQuery 如果没有值传进来,返回默认值
		// name := ctx.DefaultQuery("query", "please input something")
		name, ok := ctx.GetQuery("query")
		if !ok {
			name = "nobody"
		}
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run(":9090")
}
