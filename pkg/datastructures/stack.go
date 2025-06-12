package datastructures

import "errors"

// Stack returns a new stack with a pointer to a `node` with `num` as it's value.
func Stack(num int) stack {
	return stack{
		currentNode: &node{value: num},
	}
}

/*
Push replaces the `currentNode` with a new `node` and returns it's value.
*/
func (s *stack) Push(num int) int {
	newNode := &node{
		value: num,
		next:  s.currentNode,
	}

	s.currentNode = newNode
	return num
}

/*
Pop returns the value of the `node` referenced by `currentNode`.
It then removes that node from the sequence of `node`s.
*/
func (s *stack) Pop() (int, error) {
	if s.currentNode == nil {
		return -1, errors.New("no nodes exist in stack")
	}

	lastVal := s.currentNode.value
	s.currentNode = s.currentNode.next

	return lastVal, nil
}

// Read returns the value of the `node` that is referenced by `currentNode`
func (s *stack) Read() (int, error) {
	if s.currentNode == nil {
		return -1, errors.New("no nodes exist in stack")
	}

	return s.currentNode.value, nil
}
