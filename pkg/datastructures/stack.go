package datastructures

import "errors"

/*
stack is an implementation of a stack data structure. It contains a pointer
to a doubleLinkedNode.
*/
type stack struct {
	currentNode *doubleLinkedNode
}

// Stack returns a new stack with a pointer to a `doubleLinkedNode` with `num` as it's value.
func Stack(num int) stack {
	newNode := &doubleLinkedNode{value: num}
	return stack{
		currentNode: newNode,
	}
}

/*
Push appends a new `doubleLinkedNode` to the currentNode (as it's `next` value)
and returns it's value.
*/
func (s *stack) Push(num int) int {
	newNode := &doubleLinkedNode{
		value: num,
		prev:  s.currentNode,
	}

	s.currentNode.next = newNode
	s.currentNode = newNode
	return num
}

/*
Pop returns the value of the `doubleLinkedNode` referenced by `currentNode`.
It then removes that node from the sequence of `doubleLinkedNode`s.
*/
func (s *stack) Pop() int {
	lastVal := s.currentNode.value

	newLast := s.currentNode.prev
	s.currentNode = newLast

	return lastVal
}

// Read returns the value of the `doubleLinkedNode` that is referenced by `currentNode`
func (s *stack) Read() (int, error) {
	if s.currentNode == nil {
		return -1, errors.New("no nodes exist in stack")
	}

	return s.currentNode.value, nil
}
