package server

import (
	"github.com/g-hyoga/CP5I/server/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/recipe/:id", controller.GetRecipe)
	r.POST("/api/recipe", controller.PostRecipe)

	r.GET("/api/easy/recipe", controller.GetEasyRecipe)

	r.Run()
}
