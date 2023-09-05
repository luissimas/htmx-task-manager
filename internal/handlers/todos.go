package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/luissimas/htmx-todo/internal/database"
	. "github.com/luissimas/htmx-todo/internal/entities"
)

func Index(c *fiber.Ctx) error {
	todos := database.ListTodos()
	return c.Render("pages/index", fiber.Map{"Todos": todos}, "layouts/base")
}

func CreateTodo(c *fiber.Ctx) error {
	var todo Todo
	if err := c.BodyParser(&todo); err != nil {
		return err
	}
	todo.ID = uuid.New()
	todo.CreatedAt = time.Now()

	database.CreateTodo(todo)

	return c.Render("partials/todo-item", todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))

	database.DeleteTodo(id)

	return nil
}

func UpdateTodo(c *fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))
	todo := database.GetTodo(id)
	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	database.UpdateTodo(todo)

	return c.Render("partials/todo-item", todo)
}

func GetEditTodo(c *fiber.Ctx) error {
	id, _ := uuid.Parse(c.Params("id"))
	todo := database.GetTodo(id)

	return c.Render("partials/edit-todo-item", todo)
}
