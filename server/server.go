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

	r.GET("/api/instruct_match", controller.GetInstructMatch)

	r.Run()
}
