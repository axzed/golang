package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取请求的path(URL)参数
func main() {
	r := gin.Default()

	r.GET("/user/:name/:age", func(c *gin.Context) {
		// 获取path
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"Name": name,
			"Age":  age,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"Year":  year,
			"Month": month,
		})
	})

	r.Run(":9090")
}
