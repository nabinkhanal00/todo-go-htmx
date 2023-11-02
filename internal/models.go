package internal

import "github.com/google/uuid"

type Todo struct {
	ID        uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
}

func NewTodo(name string, completed bool) Todo {
	return Todo{
		ID:        uuid.New(),
		Name:      name,
		Completed: completed,
	}
}
