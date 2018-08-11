package todo

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	err := Create("test task.")
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
}

func TestList(t *testing.T) {
	todos, err := List()
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("result count = %d, data = %v\n", len(todos), todos)
}

func TestIsExist(t *testing.T) {
	result, err := IsExist(2)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("id exists. result = %v\n", result)

	result, err = IsExist(0)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("id not exists. result = %v\n", result)
}

func TestGet(t *testing.T) {
	todo, err := Get(2)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
	fmt.Printf("result = %v\n", todo)
}

func TestUpdate(t *testing.T) {
	err := Update(2, "test update")
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
}

func TestDelete(t *testing.T) {
	err := Delete(4)
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
}
