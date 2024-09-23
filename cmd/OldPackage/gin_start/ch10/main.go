package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 方法会将指定目录下的文件加载好，相对目录
	// 多个文件就不用这样的写法了
	// router.LoadHTMLFiles("templates/index.tmpl","templates/goods.html")

	// 加载templates下面的全部文件,包括二级目录
	router.LoadHTMLGlob("templates/**/*")


	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "擦你的",
		})
	})

	router.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"name": "你的我的",
		})
	})

	router.GET("/goods/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods/list.html", gin.H{
			"name": "商品列表",
		})
	})
	router.GET("/users/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/list.html", gin.H{
			"name": "用户列表",
		})
	})


	router.Run(":8083")
}
