package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 创建goods(商品)路由分组
	goodsGroup := router.Group("/goods")
	{
		goodsGroup.GET("", goodsList)
		goodsGroup.GET("/:id:action", goodsDetail) // url 中添加参数
		goodsGroup.POST("/add", createGoods)
	}

	//v1 := router.Group("/v1")
	//{
	//	v1.POST("/login", loginEndPoint)
	//	v1.POST("/submit", submitEndPoint)
	//	v1.POST("/read", readEndPoint)
	//}
	//
	//v2 := router.Group("/v2")
	//{
	//	v2.POST("/login", loginEndPoint)
	//	v2.POST("/submit", submitEndPoint)
	//	v2.POST("/read", readEndPoint)
	//}
	router.Run(":8083")
}

func createGoods(context *gin.Context) {

}

func goodsList(context *gin.Context) {
	context.JSONP(http.StatusOK, gin.H{
		"name": "goodslist",
	})
}

func goodsDetail(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	c.JSONP(http.StatusOK, gin.H{
		"id":     id,
		"action": action,
	})
}
