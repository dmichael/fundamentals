package stack

import (
	"fmt"

	"github.com/dmichael/fundamentals/linkedlist"
)

// Stack is a structure holding the linked list
type Stack struct {
	first *linkedlist.Node
}

// Push adds an item to the list
func (s *Stack) Push(i int) {
	next := s.first
	first := &linkedlist.Node{Item: i}
	if next != nil {
		first.Next = next
	}
	s.first = first
}

// Pop adds an item to the list
func (s *Stack) Pop() interface{} {
	head := s.first
	s.first = head.Next
	return head.Item
}

// NewStack creates a new Stack
func NewStack() *Stack {
	return &Stack{}
}

// PrintValues is a utility method so we can see whats in our stack
func (s *Stack) PrintValues() {
	current := s.first
	for current != nil {
		fmt.Println(current.Item)
		current = current.Next
	}
}
