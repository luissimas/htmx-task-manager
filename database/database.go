package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
	CREATE TABLE IF NOT EXISTS todos(
		id VARCHAR(32) PRIMARY KEY,
		text VARCHAR(255),
		done BOOLEAN,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

func Connect() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func init() {
	log.Printf("Starting database connection...")
	conn := Connect()
	conn.MustExec(schema)
	log.Printf("Successfully initialized database!")
}
