package queue

import "testing"

func TestNativeQueue(t *testing.T) {
	q := make([]int, 0)
	// push
	q = append(q, 1)
	q = append(q, 2)
	q = append(q, 3)

	value := q[0]
	// q = q[1:]

	if value != 1 {
		t.Error()
	}
	// 1 goes in first, comes out first
	// fmt.Println(q, value)
}
