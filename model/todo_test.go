package model

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	todo := Todo{}
	err := todo.Create("test task.")
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
}

func TestList(t *testing.T) {
	todo := Todo{}
	todos, err := todo.List()
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("result count = %d, data = %v\n", len(todos), todos)
}

func TestIsExist(t *testing.T) {
	todo := Todo{}
	result, err := todo.IsExist(2)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("id exists. result = %v\n", result)

	result, err = todo.IsExist(0)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("id not exists. result = %v\n", result)
}

func TestGet(t *testing.T) {
	todo := Todo{}
	data, err := todo.Get(2)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("result = %v\n", data)
}

func TestUpdate(t *testing.T) {
	todo := Todo{}
	err := todo.Update(2, "test update")
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
}

func TestDelete(t *testing.T) {
	todo := Todo{}
	err := todo.Delete(2)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
}
