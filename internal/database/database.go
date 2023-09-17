package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/luissimas/htmx-todo/internal/config"
)

var db *sqlx.DB

func init() {
	log.Printf("Starting database connection...")
	conn := connect()
	db = conn
	log.Printf("Successfully initialized database!")
}

func connect() *sqlx.DB {
	conf := config.GetConfig()
	db, err := sqlx.Connect("postgres", conf.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
