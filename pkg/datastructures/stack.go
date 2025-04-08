package datastructures

import "errors"

type stack struct {
	stack []int
}

func (s *stack) Push(num int) int {
	s.stack = append(s.stack, num)
	return num
}

func (s *stack) Pop() (int, error) {
	stackBound := len(s.stack) - 1

	if stackBound < 0 {
		return -1, errors.New("stack is empty")
	}

	last := s.stack[stackBound]
	s.stack = s.stack[:stackBound]

	return last, nil
}

func (s *stack) Read() (int, error) {
	if len(s.stack) == 0 {
		return -1, errors.New("stack is empty")
	}

	return s.stack[len(s.stack)-1], nil
}

func Stack() stack {
	return stack{
		stack: []int{},
	}
}
