package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/luissimas/htmx-todo/internal/config"
	"github.com/luissimas/htmx-todo/internal/handlers"
)

func main() {
	cfg := config.GetAPI()
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	app := fiber.New(fiber.Config{Views: initTemplateEngine()})

	initApp(app)

	log.Fatal(app.Listen(addr))
}

func initApp(app *fiber.App) {
	app.Use(recover.New())

	app.Static("/", "./web/public")

	app.Get("/", handlers.Index)
	app.Post("/todo", handlers.CreateTodo)
	app.Delete("todo/:id", handlers.DeleteTodo)
	app.Patch("todo/:id", handlers.UpdateTodo)
	app.Get("todo/:id/edit", handlers.GetEditTodo)

	app.Use(handlers.NotFound)
}

func initTemplateEngine() *html.Engine {
	engine := html.New("./web/templates", ".html")

	// Reload templates on each render (DEV only)
	engine.Reload(true)

	return engine
}
