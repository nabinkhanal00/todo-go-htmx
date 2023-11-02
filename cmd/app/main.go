package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nabinkhanal00/todo-go-htmx/internal"
	"log"
	"os"
)

func main() {

	engine := html.New("web/html", ".html")
	a := fiber.New(fiber.Config{
		Views: engine,
	})
	a.Use(cors.New())
	a.Use(logger.New())

	err := internal.SetupDB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer internal.CloseDB()

	a.Static("/static", "web/static")
	a.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	a.Listen(":" + os.Getenv("PORT"))
}
