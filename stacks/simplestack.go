package stacks

// Stack is contained in a struct to allow for mutexes and such
// for thread safety
type Stack struct {
	values []int
}

// Values exposes the values ... could have just made the struct member public
func (s *Stack) Values() []int {
	return s.values
}

// Push adds an element to the end of the stack
func (s *Stack) Push(i int) {
	s.values = append(s.values, i)
}

// Pop removes the head element of the stack
func (s *Stack) Pop() int {
	l := len(s.values)
	value := s.values[l-1]
	s.values = s.values[:l-1]
	return value
}

// NewStack returns a struct for a stack
func NewStack() *Stack {
	return &Stack{make([]int, 0)}
}
