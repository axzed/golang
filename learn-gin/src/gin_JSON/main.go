package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(ctx *gin.Context) {
		//data := map[string]interface{}{
		//	"name":    "xuewenchao",
		//	"age":     18,
		//	"message": "hello",
		//}
		// gin.H is a shortcut for map[string]interface{}
		data := gin.H{"name": "xuewenchao", "age": 18, "message": "hello gin.H"}
		ctx.JSON(http.StatusOK, data)
	})

	// use struct
	// 灵活使用tag来做定制化
	type msg struct {
		Name    string `json:"name"`
		Messaga string `json:"message"`
		Age     int    `json:"age"`
	}
	r.GET("/another_json", func(ctx *gin.Context) {
		data := msg{
			"xuewenchao",
			"hello golang",
			18,
		}
		ctx.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}
