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
