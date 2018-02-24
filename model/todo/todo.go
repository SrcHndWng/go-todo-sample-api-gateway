package todo

import (
	"github.com/SrcHndWng/go-todo-sample-api-gateway/model"
	"github.com/guregu/dynamo"
)

// Todo is todo data in DynamoDb.
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var tbl dynamo.Table

// Create registers todo.
func Create(task string) (err error) {
	tbl = table()
	var cnt int
	if cnt, err = count(); err != nil {
		return
	}
	t := Todo{ID: cnt + 1, Task: task}
	err = tbl.Put(t).Run()
	return
}

// List selects all data.
func List() (todos []Todo, err error) {
	tbl = table()
	err = tbl.Scan().All(&todos)
	return
}

// Get selects by id
func Get(id int) (todo Todo, err error) {
	tbl = table()
	err = tbl.Get("ID", id).One(&todo)
	return
}

// Update replaces data
func Update(id int, task string) error {
	tbl = table()
	return tbl.Update("ID", id).Set("Task", task).Run()
}

// Delete remove data
func Delete(id int) error {
	tbl = table()
	return tbl.Delete("ID", id).Run()
}

func count() (cnt int, err error) {
	var todos []Todo
	if todos, err = List(); err != nil {
		return
	}
	cnt = len(todos)
	return
}

func table() dynamo.Table {
	if tbl.Name() == "" {
		return model.Table("go_todos")
	}
	return tbl
}
