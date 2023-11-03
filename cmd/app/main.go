package main

import (
	"log"
	"os"

	"github.com/Masterminds/sprig/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nabinkhanal00/todo-go-htmx/internal"
)

func main() {

	engine := html.New("web/html", ".html")
	engine.AddFuncMap(sprig.FuncMap())
	a := fiber.New(fiber.Config{
		Views: engine,
	})
	a.Use(cors.New())
	a.Use(logger.New())

	err := internal.SetupDB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	a.Static("/static", "web/static")
	internal.SetupRoutes(a)
	log.Fatalln(a.Listen(":" + os.Getenv("PORT")))
}
