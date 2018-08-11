package model

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	todo := newTodoTest()
	clear(todo)
	test := func(task string) {
		err := todo.Create(task)
		if err != nil {
			t.Fatalf("error raise. %#v", err)
		}
	}
	test("test task.")
	test("test task2.")
}

func TestList(t *testing.T) {
	todo := newTodoTest()
	todos, err := todo.List()
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("result count = %d, data = %v\n", len(todos), todos)
}

func TestIsExist(t *testing.T) {
	todo := newTodoTest()
	result, err := todo.IsExist(2)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("id 2 exists. result = %v\n", result)

	result, err = todo.IsExist(0)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("id 0 not exists. result = %v\n", result)
}

func TestGet(t *testing.T) {
	todo := newTodoTest()
	data, err := todo.Get(2)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("result = %v\n", data)
}

func TestUpdate(t *testing.T) {
	todo := newTodoTest()
	err := todo.Update(2, "test update")
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
}

func TestDelete(t *testing.T) {
	todo := newTodoTest()
	err := todo.Delete(2)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
}

func newTodoTest() *Todo {
	t := Todo{}
	t.table = Table("go_todos_test")
	return &t
}

func clear(t *Todo) (err error) {
	var todos []Todo
	if todos, err = t.List(); err != nil {
		return err
	}
	for _, todo := range todos {
		t.Delete(todo.ID)
	}
	return nil
}
