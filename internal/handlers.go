package internal

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	todo, err := createTodo(name)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	count, err := getTotalCount()
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	renderTemplate(c, "TotalCount", fiber.Map{"SwapOOB": true, "Count": count})
	renderTemplate(c, "Todo", fiber.Map{"SwapOOB": true, "Todo": todo})
	renderTemplate(c, "Form", nil)
	return c.SendStatus(fiber.StatusCreated)
}

func HandlerGetTodo(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	todo, err := getTodo(id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	renderTemplate(c, "Todo", fiber.Map{"SwapOOB": false, "Todo": todo, "Editing": true})
	return c.SendStatus(fiber.StatusOK)

}

func HandlerEditTodo(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	name := c.FormValue("name")

	if name == "" {
		todo, err := getTodo(id)
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		renderTemplate(c, "Todo", fiber.Map{"SwapOOB": false, "Todo": todo, "Editing": false})
		return c.SendStatus(fiber.StatusBadRequest)
	}
	todo, err := editTodo(id, name)
	if err != nil {

		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	renderTemplate(c, "Todo", fiber.Map{"SwapOOB": false, "Todo": todo, "Editing": false})
	return c.SendStatus(fiber.StatusOK)
}

func HandlerToggleTodo(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}
	todo, err := toggleTodo(id)
	if err != nil {
		return err
	}
	count, err := getCompletedCount()
	if err != nil {
		return err
	}
	renderTemplate(c, "CompletedCount", fiber.Map{"SwapOOB": true, "Count": count})
	renderTemplate(c, "Todo", fiber.Map{"ID": todo.ID, "Name": todo.Name, "Completed": todo.Completed})
	return c.SendStatus(fiber.StatusAccepted)
}
func HandlerDeleteTodo(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}
	_, err = deleteTodo(id)

	completedCount, err := getCompletedCount()
	if err != nil {
		return err
	}
	totalCount, err := getTotalCount()

	if err != nil {
		return err
	}
	renderTemplate(c, "TotalCount", fiber.Map{"SwapOOB": true, "Count": totalCount})
	renderTemplate(c, "CompletedCount", fiber.Map{"SwapOOB": true, "Count": completedCount})
	return c.SendStatus(fiber.StatusAccepted)
}
