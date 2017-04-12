package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(6)
	s.PrintValues()
}

func TestStackPop(t *testing.T) {
	s := NewStack()
	s.Push(5)
	s.Push(4)
	s.Push(3)

	v := s.Pop()
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
