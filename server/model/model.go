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
