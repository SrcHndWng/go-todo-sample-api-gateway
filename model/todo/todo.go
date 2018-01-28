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

// Create registers todo.
func Create(task string) error {
	dy := model.DynamoDb()
	tbl := table(dy)
	cnt, err := count(tbl)
	if err != nil {
		return err
	}
	t := Todo{ID: cnt + 1, Task: task}
	return tbl.Put(t).Run()
}

func count(table dynamo.Table) (int, error) {
	var todos []Todo
	err := table.Scan().All(&todos)
	if err != nil {
		return 0, err
	}
	return len(todos), nil
}

func table(dy *dynamo.DB) dynamo.Table {
	return dy.Table("go_todos")
}
