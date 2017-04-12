package stack

import (
	"fmt"
	"testing"
)

func TestNativeStack(t *testing.T) {
	s := make([]int, 0)

	// push
	s = append(s, 8)
	s = append(s, 9)
	fmt.Println(s)

	// pop
	value := s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Println(value, s)
}
