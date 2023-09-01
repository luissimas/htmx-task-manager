package database

import (
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/luissimas/htmx-todo/todos"
)

func ListTodos(db *sqlx.DB) (todos []todos.Todo) {
	err := db.Select(&todos, "SELECT id, text, done FROM todos ORDER BY created_at DESC")
	if err != nil {
		log.Fatal(err)
	}
	return todos
}

func GetTodo(db *sqlx.DB, id uuid.UUID) (todo todos.Todo) {
	err := db.Get(&todo, "SELECT id, text, done FROM todos WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	return todo
}

func CreateTodo(db *sqlx.DB, todo todos.Todo) {
	_, err := db.NamedExec("INSERT INTO todos (id, text, done) VALUES (:id, :text, :done)", todo)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateTodo(db *sqlx.DB, todo todos.Todo) {
	_, err := db.NamedExec("UPDATE todos SET text=:text, done=:done WHERE id=:id", todo)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteTodo(db *sqlx.DB, id uuid.UUID) {
	_, err := db.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
}
