package internal

import "github.com/gofiber/fiber/v2"

func SetupRoutes(a *fiber.App) {
	a.Put("/:id/toggle", HandlerToggleTodo)
	a.Put("/:id/edit", HandlerEditTodo)
	a.Get("/:id", HandlerGetTodo)
	a.Delete("/:id", HandlerDeleteTodo)
	a.Get("/", HandlerGetTodos)
	a.Post("/", HandlerCreateTodo)
}
