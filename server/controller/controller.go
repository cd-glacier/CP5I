package controller

import (
	"github.com/g-hyoga/CP5I/server/model"
	"github.com/gin-gonic/gin"
)

func GetRecipe(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": model.Recipe{},
	})
}
