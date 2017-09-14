package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/g-hyoga/CP5I/server/model"
	"github.com/g-hyoga/CP5I/server/score"
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

func GetEasyRecipes(c *gin.Context) {
	err := db.Connect()
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	recipes := []model.Recipe{}
	food := c.Query("food")
	kithechware := c.Query("kithechware")
	recipes, err = db.GetEasyRecipes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if food != "" || kithechware != "" {
		recipes, _ = Filter(strings.Split(food, ","), strings.Split(kithechware, ","), recipes)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": recipes,
	})
}

func Contains(array []string, target string) bool {
	for _, e := range array {
		if strings.Contains(target, e) {
			return true
		}
	}
	return false
}

func Filter(food []string, kithechwares []string, recipes []model.Recipe) ([]model.Recipe, error) {
	result := []model.Recipe{}
	for _, recipe := range recipes {
		for _, ingredient := range recipe.Ingredients {
			if Contains(food, ingredient.Name) {
				result = append(result, recipe)
				break
			}
		}
	}

	return result, nil
}

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

	//scoringa
	recipe.Difficulty, _ = score.Score(recipe.Method)

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
