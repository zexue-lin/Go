package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 创建goods(商品)路由分组
	goodsGroup := router.Group("/goods")
	{
		goodsGroup.GET("", goodsList)
		goodsGroup.GET("/:id/:action", goodsDetail) // url 中添加参数
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

/*
	针对于方法 goodsDetail
	此时的utl地址是 http://127.0.0.1/goods/1/delete
	获取url中的参数 。GET
	比如获取商品的相信信息，根据商品id来查询

	但是这种匹配模式可以把你url中的list也匹配出来，
	比如 http://127.0.0.1:8083/goods/list/delete
	json返回{
		action: "delete",
		id: "list"
	}

	这样明显就不符合规范
*/
