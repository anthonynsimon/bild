package util

//A utility stack implementation for arbitrary data types with Push(), Pop() and Len() functions
type Stack struct {
	top  *stackElement
	size int
}

type stackElement struct {
	value interface{}
	next  *stackElement
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &stackElement{value, s.top}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		value := s.top.value
		s.top = s.top.next
		s.size--
		return value
	}
	return nil
}
