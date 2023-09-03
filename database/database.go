package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/luissimas/htmx-todo/config"
)

var schema = `
	CREATE TABLE IF NOT EXISTS todos(
		id UUID PRIMARY KEY,
		text VARCHAR(255),
		done BOOLEAN,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`
var db *sqlx.DB

func init() {
	log.Printf("Starting database connection...")
	conn := connect()
	conn.MustExec(schema)
	db = conn
	log.Printf("Successfully initialized database!")
}

func connect() *sqlx.DB {
	conf := config.GetDB()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Database)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
