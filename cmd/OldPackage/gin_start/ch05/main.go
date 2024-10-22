package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/welcome", welcome)
	router.POST("/form_post", formPost)
	router.POST("/post", getPost)
	router.Run(":8083")
}

// post 和 get 混合获取参数
func getPost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0") // GET 方式
	name := c.PostForm("name")          // post方式
	message := c.DefaultPostForm("message", "信息")

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}

// Post
func formPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"nick":    nick,
	})
}

// Get
func welcome(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "tom") // 取不到默认值是tom
	lastName := c.Query("lastname")                 // 没有默认值

	c.JSON(http.StatusOK, gin.H{
		"first_name": firstName,
		"last_name":  lastName,
	})
	// 地址是 http://127.0.0.1:8083/welcome?firstname=tom&lastname=jerry
}
