package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func renderTemplate(c *fiber.Ctx, template string, data interface{}) error {
	return c.App().Config().Views.(*html.Engine).Templates.ExecuteTemplate(c.Response().BodyWriter(), template, data)
}
