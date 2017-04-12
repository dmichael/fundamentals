package stacks

import (
	"reflect"
	"testing"
)

var fixture = []int{5, 4, 3}

func setup() *Stack {
	s := NewStack()
	for _, val := range fixture {
		s.Push(val)
	}
	return s
}

// Stack em up
func TestPush(t *testing.T) {
	s := setup()
	if !reflect.DeepEqual(s.Values(), fixture) {
		t.Error("Stack is not populating correctly  ")
	}
}

// Last in, first out
func TestPop(t *testing.T) {
	s := setup()

	v := s.Pop()
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
