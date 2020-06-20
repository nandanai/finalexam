package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	createTb := `CREATE TABLE if NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
		);`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table customers", err)
	}
}

func Conn() *sql.DB {
	return db
}

func DeleteByID(id string) error {

	stmt, err := Conn().Prepare("DELETE FROM customer WHERE id = $1")
	if err != nil {
		return fmt.Errorf("can't prepare delete statement: %w", err)
	}

	if _, err := stmt.Exec(id); err != nil {
		return fmt.Errorf("can't execute delete statment: %w", err)
	}

	return nil
}
