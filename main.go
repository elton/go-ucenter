package main

import "github.com/gin-gonic/gin"

func main() {
	// 初始化引擎
	r := gin.Default()

	// 注册路由器
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 运行，默认开启8080端口，也可以自定义
	r.Run() // listen and serve on 0.0.0.0:8080
}
