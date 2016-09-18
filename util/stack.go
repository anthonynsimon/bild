package util

// Stack implementation for arbitrary data types with Push(), Pop() and Len() functions
type Stack struct {
	top  *stackElement
	size int
}

type stackElement struct {
	value interface{}
	next  *stackElement
}

// Len returns the size of stack
func (s *Stack) Len() int {
	return s.size
}

// Push a new value onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &stackElement{value, s.top}
	s.size++
}

// Pop the most recently pushed value from the stack
func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		value := s.top.value
		s.top = s.top.next
		s.size--
		return value
	}
	return nil
}
