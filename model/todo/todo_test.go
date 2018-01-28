package todo

import (
	"testing"
)

func TestCreate(t *testing.T) {
	err := Create("test task.")
	if err != nil {
		t.Fatalf("error raise. %#v", err)
	}
}
