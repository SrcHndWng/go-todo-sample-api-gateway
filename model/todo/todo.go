package todo

import (
	"../../model"
	"github.com/guregu/dynamo"
)

// Todo is todo data in DynamoDb.
type Todo struct {
	ID   int
	Task string
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
