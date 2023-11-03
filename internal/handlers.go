package internal

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func HandlerGetTodos(c *fiber.Ctx) error {
	todos, err := getTodos()
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	completedCount := 0
	for _, el := range todos {
		if el.Completed {
			completedCount++
		}
	}
	return c.Status(fiber.StatusOK).Render("index", fiber.Map{"Todos": todos, "TotalCount": len(todos), "CompletedCount": completedCount})

}

func HandlerCreateTodo(c *fiber.Ctx) error {
	name := c.FormValue("name")
	if name == "" {
		return c.Render("Form", nil)
	}
	_, err := createTodo(name)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusCreated).Render("Form", nil)
}

func HandlerEditTodo(c *fiber.Ctx) error {
	return nil
}

func HandlerDeleteTodo(c *fiber.Ctx) error {
	return nil
}
