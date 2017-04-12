package stack

// SimpleStack is contained in a struct to allow for mutexes and such
// for thread safety
type SimpleStack struct {
	values []int
}

// Values exposes the values ... could have just made the struct member public
func (s *SimpleStack) Values() []int {
	return s.values
}

// Push adds an element to the end of the stack
func (s *SimpleStack) Push(i int) {
	s.values = append(s.values, i)
}

// Pop removes the head element of the stack
func (s *SimpleStack) Pop() int {
	l := len(s.values)
	value := s.values[l-1]
	s.values = s.values[:l-1]
	return value
}

// NewSimpleStack returns a struct for a stack
func NewSimpleStack() *SimpleStack {
	return &SimpleStack{make([]int, 0)}
}
