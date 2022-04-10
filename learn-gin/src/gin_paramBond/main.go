package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// shouldBind 反射取字段
type UserInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pwd"`
}

// 参数绑定
func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		// 请求数据量太大的时候,这样太麻烦了
		//username := c.Query("username")
		//password := c.Query("password")
		//u := UserInfo{
		//	username: username,
		//	password: password,
		//}
		var u UserInfo          // 声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u) // 引用传递才能修改值
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
		fmt.Printf("%#v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"Message": "ok",
		})
	})

	r.Run(":9090")
}
