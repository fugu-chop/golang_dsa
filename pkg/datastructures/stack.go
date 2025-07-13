package datastructures

import "errors"

/*
stack is an implementation of a stack data structure. It contains a pointer
to a `node`.
*/
type stack[T any] struct {
	currentNode *node[T]
}

// Stack returns a new stack with a pointer to a `node` with `num` as it's value.
func Stack[T any](val T) stack[T] {
	return stack[T]{
		currentNode: &node[T]{value: val},
	}
}

/*
Push replaces the `currentNode` with a new `node` and returns it's value.
*/
func (s *stack[T]) Push(val T) T {
	newNode := &node[T]{
		value: val,
		next:  s.currentNode,
	}

	s.currentNode = newNode
	return val
}

/*
Pop returns the value of the `node` referenced by `currentNode`.
It then removes that node from the sequence of `node`s.
*/
func (s *stack[T]) Pop() (T, error) {
	if s.currentNode == nil {
		return *new(T), errors.New("no nodes exist in stack")
	}

	lastVal := s.currentNode.value
	s.currentNode = s.currentNode.next

	return lastVal, nil
}

// Read returns the value of the `node` that is referenced by `currentNode`
func (s *stack[T]) Read() (T, error) {
	if s.currentNode == nil {
		return *new(T), errors.New("no nodes exist in stack")
	}

	return s.currentNode.value, nil
}
