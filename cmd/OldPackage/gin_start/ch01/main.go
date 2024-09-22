package main

import "github.com/gin-gonic/gin"

//func pong(c *gin.Context) {
//	c.JSON(http.StatusOK, map[string]string{
//		"message": "pong",
//	})
//}

// 跟上面一样，H 其实就是一个map
//
//	func pong(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"message": "pong",
//		})
//	}
func main() {
	//r := gin.Default()
	//r.GET("/ping", pong)
	//r.Run(":8083") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// 使用默认中间件创建一个gin路由器
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/someGet", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)

	router.Run()
}
