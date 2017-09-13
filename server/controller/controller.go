package controller

import (
	"net/http"
	"strconv"

	"github.com/g-hyoga/CP5I/server/model"
	"github.com/gin-gonic/gin"
)

var db model.DB

func GetRecipe(c *gin.Context) {
	err := db.Connect()
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	recipe, err := db.GetRecipe(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": recipe,
	})
}

func GetEasyRecipe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": model.Recipe{},
	})
}

// scoring まだしていない
func PostRecipe(c *gin.Context) {
	var recipe model.Recipe
	c.BindJSON(&recipe)

	err := db.Connect()
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	// duplicated data?
	id, err := db.GetRecipeID(recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	if id >= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "It is duplicated data",
		})
		return
	}

	//scoring
	recipe.Difficulty = 10

	// insert
	err = db.InsertRecipe(recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	recipe.ID, err = db.GetRecipeID(recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	err = db.InsertIngredients(recipe.ID, recipe.Ingredients)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	err = db.InsertMethod(recipe.ID, recipe.Method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success",
		"data":   recipe,
	})
}
