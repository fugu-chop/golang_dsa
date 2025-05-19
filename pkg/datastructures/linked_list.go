package datastructures

import "fmt"

type LinkedList struct {
	node *node
}

type node struct {
	value *int
	next  *node
}

func New(val int) *LinkedList {
	return &LinkedList{
		node: &node{
			value: &val,
		},
	}
}

func (l *LinkedList) ReadAt(idx int) (int, error) {
	n := l.node
	for range idx {
		if n.next == nil {
			return -1, fmt.Errorf("node at index %d does not exist", idx)
		}
		n = n.next
	}

	return *n.value, nil
}

func (l *LinkedList) Search(value int) (int, error) {
	idx := 0
	n := l.node

	for n.next != nil {
		if *n.value == value {
			return idx, nil
		}
		n = n.next
		idx++
	}

	return -1, fmt.Errorf("value '%d' does not exist in linked list", value)
}
