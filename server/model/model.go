package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	db *sql.DB
}

func (db *DB) Connect() error {
	var err error
	db.db, err = sql.Open("mysql", "root:password@/cookpad")
	return err
}

func (db *DB) Close() error {
	err := db.db.Close()
	return err
}

func (db *DB) InsertRecipe(r Recipe) error {
	_, err := db.db.Exec("INSERT INTO `recipe` (id, name, time, producer_id, difficulty) VALUES(?, ?, ?, ?, ?);", 0, r.Name, r.Time, r.ProducerID, r.Difficulty)
	return err
}

func (db *DB) InsertIngredients(recipeID int, ingredients []Ingredient) error {
	var err error
	for _, ingredient := range ingredients {
		_, err = db.db.Exec("INSERT INTO `ingredient` (id, recipe_id, name, quantity) VALUES(?, ?, ?, ?);", 0, recipeID, ingredient.Name, ingredient.Quantity)
	}

	return err
}

func (db *DB) InsertMethod(recipeID int, ms []Method) error {
	var err error
	for i, m := range ms {
		_, err = db.db.Exec("INSERT INTO `method` (id, recipe_id, method_order, content) VALUES(?, ?, ?, ?);", 0, recipeID, i, m.Content)
	}

	return err
}

// first data is -1
func (db *DB) GetRecipeID(r Recipe) (int, error) {
	sql := "select * from `recipe` where name=? and time=? and producer_id=? and difficulty=?;"
	rows, err := db.db.Query(sql, r.Name, r.Time, r.ProducerID, r.Difficulty)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	recipes := []Recipe{}
	recipes, err = scanRecipe(rows)
	if err != nil {
		return -1, err
	}
	if len(recipes) == 0 {
		return -1, nil
	}

	return recipes[0].ID, nil
}

func (db *DB) GetRecipe(id int) (Recipe, error) {
	var recipe Recipe

	sql := "select * from `recipe` where id=?;"
	rows, err := db.db.Query(sql, id)
	if err != nil {
		return recipe, err
	}
	defer rows.Close()
	recipes := []Recipe{}
	recipes, err = scanRecipe(rows)
	if err != nil {
		return recipe, err
	}

	if len(recipes) == 0 {
		return recipe, nil
	}

	recipe = recipes[0]
	recipe.Ingredients, err = db.GetIngredients(id)
	if err != nil {
		return recipe, err
	}
	recipe.Method, err = db.GetMethod(id)
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}

func (db *DB) GetEasyRecipes() ([]Recipe, error) {
	recipes := []Recipe{}

	sql := "select * from `recipe` order by difficulty desc limit 10;"
	rows, err := db.db.Query(sql)
	if err != nil {
		return recipes, err
	}
	defer rows.Close()
	recipes, err = scanRecipe(rows)
	if err != nil {
		return recipes, err
	}

	if len(recipes) == 0 {
		return recipes, nil
	}

	for i, recipe := range recipes {
		id, err := db.GetRecipeID(recipe)
		if err != nil {
			return recipes, err
		}

		recipes[i].Ingredients, err = db.GetIngredients(id)
		if err != nil {
			return recipes, err
		}
		recipes[i].Method, err = db.GetMethod(id)
		if err != nil {
			return recipes, err
		}
	}

	return recipes, nil
}

func (db *DB) GetIngredients(recipeID int) ([]Ingredient, error) {
	sql := "select * from `ingredient` where recipe_id=?;"
	rows, err := db.db.Query(sql, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ingredients := []Ingredient{}
	ingredients, err = scanIngredient(rows)
	if err != nil {
		return nil, err
	}
	return ingredients, nil
}

func (db *DB) GetMethod(recipeID int) ([]Method, error) {
	sql := "select * from `method` where recipe_id=?;"
	rows, err := db.db.Query(sql, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := []Method{}
	ms, err = scanMethod(rows)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

func scanRecipe(rows *sql.Rows) ([]Recipe, error) {
	recipes := []Recipe{}
	var err error
	for rows.Next() {
		var r Recipe
		if err = rows.Scan(&r.ID, &r.Name, &r.Time, &r.ProducerID, &r.Difficulty); err != nil {
			return recipes, err
		}
		recipes = append(recipes, r)
	}
	return recipes, err
}

func scanIngredient(rows *sql.Rows) ([]Ingredient, error) {
	ingredients := []Ingredient{}
	var err error
	for rows.Next() {
		var ingredient Ingredient
		if err = rows.Scan(&ingredient.ID, &ingredient.RecipeID, &ingredient.Name, &ingredient.Quantity); err != nil {
			return ingredients, err
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, err
}

func scanMethod(rows *sql.Rows) ([]Method, error) {
	ms := []Method{}
	var err error
	for rows.Next() {
		var m Method
		if err = rows.Scan(&m.ID, &m.RecipeID, &m.Order, &m.Content); err != nil {
			return ms, err
		}
		ms = append(ms, m)
	}
	return ms, err
}
