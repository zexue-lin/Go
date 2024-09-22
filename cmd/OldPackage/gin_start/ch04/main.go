package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID   int    `uri:"id" binding:"required"` // 添加约束
	Name string `uri:"name" binding:"required"`
}

func main() {
	router := gin.Default()

	router.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.Status(404)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": person.Name,
			"id":   person.ID,
		})
	})
	err := router.Run(":8084")
	if err != nil {
		return
	}
}
