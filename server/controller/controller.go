package controller

import (
	"net/http"

	"github.com/g-hyoga/CP5I/server/model"
	"github.com/gin-gonic/gin"
)

var db model.DB

func GetRecipe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": model.Recipe{},
	})
}

func PostRecipe(c *gin.Context) {
	var recipe model.Recipe
	c.BindJSON(&recipe)

	err := db.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	err = db.InsertRecipe(recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	err = db.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}
