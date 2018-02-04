package todos

import (
	"../../model/todo"
)

// Create controller.
func Create(task string) error {
	return todo.Create(task)
}

// List controller.
func List() ([]todo.Todo, error) {
	todos, err := todo.List()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// Get controller.
func Get(id int) (todo.Todo, error) {
	return todo.Get(id)
}
