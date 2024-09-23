package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义一个中间件
func MyLogger() gin.HandlerFunc { // 返回一个 gin.HandlerFunc 类型的值。
	return func(c *gin.Context) {
		t := time.Now()

		// 让原本该执行的逻辑继续执行
		c.Next() //调用下一个中间件或路由处理器。这是确保中间件链中其他部分能够正常执行的关键步骤。

		end := time.Since(t)
		fmt.Printf("耗时:%v\n", end) // 打印运行时间

		status := c.Writer.Status()
		fmt.Printf("状态是:%v\n", status)
	}
}

func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			if k == "X-Token" { // 这里一定要大写 X T
				token = v[0] // 因为Header的值 是一个字符串切片。源码 -> type Header map[string][]string
			} else {
				fmt.Println(k, v)
			}
			fmt.Println(k, v, token)
		}
		if token != "lattiex" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			// return // 还是会执行后面的逻辑，执行c.Next() 打印pong

			// 这个就可以阻止c.Next()。中止执行
			c.Abort()
		}
		c.Next()
	}
}

func main() {
	// router := gin.Default() // 启动的时候会自动启动Logger()和recovery() 这两个中间件

	// router := gin.New()
	// 手动启动 使用Logger中间件和Recovery中间件
	// router.Use(MyLogger())

	router := gin.Default()

	router.Use(TokenRequired())

	// authoried := router.Group("/goods")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	_ = router.Run(":8083")

}
