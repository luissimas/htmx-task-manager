package database

import (
	"log"

	"github.com/google/uuid"
	. "github.com/luissimas/htmx-todo/entities"
)

func ListTodos() (todos []Todo) {
	err := db.Select(&todos, "SELECT id, text, done FROM todos ORDER BY created_at DESC")
	if err != nil {
		log.Fatal(err)
	}
	return todos
}

func GetTodo(id uuid.UUID) (todo Todo) {
	err := db.Get(&todo, "SELECT id, text, done FROM todos WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	return todo
}

func CreateTodo(todo Todo) {
	_, err := db.NamedExec("INSERT INTO todos (id, text, done) VALUES (:id, :text, :done)", todo)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateTodo(todo Todo) {
	_, err := db.NamedExec("UPDATE todos SET text=:text, done=:done WHERE id=:id", todo)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteTodo(id uuid.UUID) {
	_, err := db.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
}
