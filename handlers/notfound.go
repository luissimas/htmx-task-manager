package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).Render("pages/errors/404", nil, "layouts/base")
}
