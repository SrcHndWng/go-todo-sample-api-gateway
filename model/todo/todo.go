package todo

import (
	"../../model"
	"github.com/guregu/dynamo"
)

// Todo is todo data in DynamoDb.
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var tbl dynamo.Table

// Create registers todo.
func Create(task string) error {
	tbl = table()
	cnt, err := count()
	if err != nil {
		return err
	}
	t := Todo{ID: cnt + 1, Task: task}
	return tbl.Put(t).Run()
}

// List selects all data.
func List() ([]Todo, error) {
	tbl = table()
	var todos []Todo
	err := tbl.Scan().All(&todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// Get selects by id
func Get(id int) (Todo, error) {
	tbl = table()
	var todo Todo
	err := tbl.Get("ID", id).One(&todo)
	return todo, err
}

// Update replaces data
func Update(id int, task string) error {
	tbl = table()
	err := tbl.Update("ID", id).Set("Task", task).Run()
	if err != nil {
		return err
	}
	return nil
}

// Delete remove data
func Delete(id int) error {
	tbl = table()
	err := tbl.Delete("ID", id).Run()
	if err != nil {
		return err
	}
	return nil
}

func count() (int, error) {
	todos, err := List()
	if err != nil {
		return 0, err
	}
	return len(todos), nil
}

func table() dynamo.Table {
	if tbl.Name() == "" {
		return model.Table("go_todos")
	}
	return tbl
}
