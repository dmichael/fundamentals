package stack

import (
	"fmt"

	"github.com/dmichael/fundamentals/linkedlist"
)

// Stack is a structure holding the linked list
type Stack struct {
	Head *linkedlist.Node
}

// Push adds an item to the list
func (s *Stack) Push(i int) {
	head := s.Head
	first := &linkedlist.Node{Item: i}
	first.Next = head
	s.Head = first
}

// Pop adds an item to the list
func (s *Stack) Pop() interface{} {
	head := s.Head
	s.Head = head.Next
	return head.Item
}

// NewStack creates a new Stack
func NewStack() *Stack {
	return &Stack{}
}

// PrintValues is a utility method so we can see whats in our stack
func (s *Stack) PrintValues() {
	current := s.Head
	for current != nil {
		fmt.Println(current.Item)
		current = current.Next
	}
}
