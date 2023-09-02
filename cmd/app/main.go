package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/luissimas/htmx-todo/handlers"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: initTemplateEngine(),
	})
	initApp(app)

	log.Fatal(app.Listen(":3000"))
}

func initApp(app *fiber.App) {
	app.Use(recover.New())

	app.Static("/public", "./public")

	app.Get("/", handlers.Index)
	app.Post("/todo", handlers.CreateTodo)
	app.Delete("todo/:id", handlers.DeleteTodo)
	app.Patch("todo/:id", handlers.UpdateTodo)
	app.Get("todo/:id/edit", handlers.GetEditTodo)

	app.Use(handlers.NotFound)
}

func initTemplateEngine() *html.Engine {
	engine := html.New("./views", ".html")

	// Reload templates on each render (DEV only)
	engine.Reload(true)

	return engine
}
