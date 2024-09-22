package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/welcome", welcome)
	router.POST("/form_post", formpost)
	router.POST("/post", getPost)
	router.Run(":8083")
}

func getPost(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.DefaultPostForm("message", "信息")

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}

func formpost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"nick":    nick,
	})
}

func welcome(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "tom") // 取不到默认值是tom
	lastName := c.DefaultQuery("lastname", "last")  // 没有默认值

	c.JSON(http.StatusOK, gin.H{
		"first_name": firstName,
		"last_name":  lastName,
	})
}
