package internal

import (
	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID
	Name      string
	Completed bool
}

func getTodos() ([]*Todo, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, completed from todos;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []*Todo{}
	for rows.Next() {
		item := &Todo{}
		err := rows.Scan(&item.ID, &item.Name, &item.Completed)
		if err != nil {
			return todos, err
		}
		todos = append(todos, item)
	}
	return todos, nil
}
func getTodo(id uuid.UUID) (*Todo, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	todo := &Todo{}
	err = db.QueryRow("SELECT id, name, completed FROM todos WHERE id = (?);", id).Scan(&todo.ID, &todo.Name, &todo.Completed)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func editTodo(id uuid.UUID, name string, completed bool) (*Todo, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	todo := &Todo{}
	err = db.QueryRow("UPDATE todos SET name = (?), completed = (?) where id = (?) returning id, name, completed;", name, completed, id).Scan(&todo.ID, &todo.Name, &todo.Completed)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func createTodo(name string) (*Todo, error) {

	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	id := uuid.New()

	_, err = db.Exec(`INSERT INTO todos (id, name) VALUES (?, ?);`, id, name)
	if err != nil {
		return nil, err
	}
	return &Todo{ID: id, Name: name, Completed: false}, nil
}

func deleteTodo(id uuid.UUID) (*Todo, error) {

	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	todo := &Todo{ID: id}
	err = db.QueryRow("DELETE FROM todos WHERE id = (?) returning name, completed;", id).Scan(&todo.Name, &todo.Completed)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
