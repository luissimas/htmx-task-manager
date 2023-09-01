package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/google/uuid"
	"github.com/luissimas/htmx-todo/database"
	"github.com/luissimas/htmx-todo/todos"
)

func init() {
	database.Connect()
}

func main() {
	engine := html.New("./views", ".html")

	// Reload templates on each render (DEV only)
	engine.Reload(true)

	db := database.Connect()
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		todos := database.ListTodos(db)
		return c.Render("index", fiber.Map{"Todos": todos})
	})
	app.Post("/todo", func(c *fiber.Ctx) error {
		id := uuid.New()
		text := c.FormValue("todo")
		todo := todos.Todo{Id: id, Text: text, Done: false, CreatedAt: time.Now()}

		database.CreateTodo(db, todo)
		log.Printf("Created new todo item: %v", todo)

		return c.Render("todo", todo)
	})
	app.Delete("todo/:id", func(c *fiber.Ctx) error {
		id, _ := uuid.Parse(c.Params("id"))

		database.DeleteTodo(db, id)
		log.Printf("Deleted todo item: %v", id)

		return nil
	})
	app.Patch("todo/:id", func(c *fiber.Ctx) error {
		id, _ := uuid.Parse(c.Params("id"))
		done, _ := strconv.ParseBool(c.FormValue("done"))
		todo := database.GetTodo(db, id)
		todo.Done = done

		database.UpdateTodo(db, todo)
		log.Printf("Updated todo item: %v", id)

		return c.Render("todo", todo)
	})

	log.Fatal(app.Listen(":3000"))
}
