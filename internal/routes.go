package internal

import "github.com/gofiber/fiber/v2"

func SetupRoutes(a *fiber.App) {
	a.Post("/", HandlerCreateTodo)
	a.Put("/:id", HandlerEditTodo)
	a.Delete("/:id", HandlerDeleteTodo)
	a.Get("/", HandlerGetTodos)
}
