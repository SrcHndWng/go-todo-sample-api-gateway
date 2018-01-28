package todos

import (
	"fmt"

	"../../model/todo"
)

// Create controller.
func Create(task string) error {
	fmt.Println("todos create!! task = " + task)
	return todo.Create(task)
}

// List controller.
func List() {
	fmt.Println("todos list!!")
}
