package model

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

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
