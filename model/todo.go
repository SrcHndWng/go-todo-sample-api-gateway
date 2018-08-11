package model

import (
	"github.com/guregu/dynamo"
)

// Todo is todo data in DynamoDb.
type Todo struct {
	ID    int    `json:"id"`
	Task  string `json:"task"`
	table dynamo.Table
}

// NewTodo creates Todo Instance.
func NewTodo() *Todo {
	t := Todo{}
	t.table = Table("go_todos")
	return &t
}

// Create registers todo.
func (t *Todo) Create(task string) (err error) {
	var cnt int
	if cnt, err = t.count(); err != nil {
		return
	}
	data := Todo{ID: cnt + 1, Task: task}
	err = t.table.Put(data).Run()
	return
}

// List selects all data.
func (t *Todo) List() (todos []Todo, err error) {
	err = t.table.Scan().All(&todos)
	return
}

// IsExist returns whether id exists.
func (t *Todo) IsExist(id int) (result bool, err error) {
	cnt, err := t.table.Get("ID", id).Count()
	result = (cnt == 1)
	return
}

// Get selects by id
func (t *Todo) Get(id int) (todo Todo, err error) {
	err = t.table.Get("ID", id).One(&todo)
	return
}

// Update replaces data
func (t *Todo) Update(id int, task string) error {
	return t.table.Update("ID", id).Set("Task", task).Run()
}

// Delete remove data
func (t *Todo) Delete(id int) error {
	return t.table.Delete("ID", id).Run()
}

func (t *Todo) count() (cnt int, err error) {
	var todos []Todo
	if todos, err = t.List(); err != nil {
		return
	}
	cnt = len(todos)
	return
}
