package queue

import (
	"fmt"

	"github.com/dmichael/fundamentals/linkedlist"
)

// Queue needs comments
type Queue struct {
	head *linkedlist.Node
	tail *linkedlist.Node
}

// IsEmpty needs comments
func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

// Enqueue needs comments
func (q *Queue) Enqueue(item interface{}) {
	oldtail := q.tail
	q.tail = &linkedlist.Node{Item: item}
	if q.IsEmpty() {
		q.head = q.tail
	} else {
		oldtail.Next = q.tail
	}
}

// Dequeue needs comments
func (q *Queue) Dequeue() interface{} {
	node := q.head
	q.head = node.Next
	if q.IsEmpty() {
		q.tail = nil
	}
	return node.Item
}

// PrintValues needs comments
func (q *Queue) PrintValues() {
	current := q.head
	for current != nil {
		fmt.Println(current.Item)
		current = current.Next
	}
}

// NewQueue needs comments
func NewQueue() *Queue {
	return &Queue{}
}
