package queue

import "testing"

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	// fmt.Printf("%#v\n", q)
	// q.PrintValues()
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	value := q.Dequeue()
	if value != 1 {
		t.Error("Nope")
	}

	value = q.Dequeue()
	if value != 2 {
		t.Error("Nope")
	}
	// q.PrintValues()
}
