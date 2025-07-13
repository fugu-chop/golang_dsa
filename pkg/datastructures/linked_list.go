package datastructures

import (
	"fmt"
	"reflect"
)

/*
linkedList is an implementation of a single Linked List
data structure. It contains a pointer to a `node` type.
*/
type linkedList[T any] struct {
	node *node[T]
}

// NewLinkedList returns a pointer to a new linkedList.
func LinkedList[T any](val T) *linkedList[T] {
	return &linkedList[T]{
		node: &node[T]{
			value: val,
		},
	}
}

/*
ReadAt follows the list of nodes in the Linked List until it
reaches the node at index `idx`.

It returns an error if there is no node at index `idx`.
*/
func (l *linkedList[T]) ReadAt(idx int) (T, error) {
	n := l.node
	for range idx {
		if n.next == nil {
			return *new(T), fmt.Errorf("node at index %d does not exist", idx)
		}
		n = n.next
	}

	return n.value, nil
}

/*
IndexOf searches the Linked List for a node that has a value of `value`,
returning the 'index' of the node that contains the value.

It returns an error if the `value` does not exist within any of the nodes in
the Linked List.
*/
func (l *linkedList[T]) IndexOf(value T) (int, error) {
	idx := 0
	n := l.node

	for n != nil {
		if reflect.DeepEqual(n.value, value) {
			return idx, nil
		}
		n = n.next
		idx++
	}

	return -1, fmt.Errorf("value '%v' does not exist in linked list", value)
}

/*
InsertAt inserts a node containing a value of `value` at the `idx`th
node of the Linked List.

It returns an error if there is no complete chain of nodes leading to
the node at index `idx`.
*/
func (l *linkedList[T]) InsertAt(idx int, val T) error {
	newNode := &node[T]{
		value: val,
	}

	if idx == 0 {
		newNode.next = l.node.next
		l.node = newNode
		return nil
	}

	n := l.node

	for range idx - 1 {
		n = n.next
		if n == nil {
			return fmt.Errorf("cannot insert node at index %d due to broken chain", idx)
		}
	}

	newNode.next = n.next
	n.next = newNode

	return nil
}

/*
DeleteAt removes a node at the `idx`th node of a Linked List.
It returns an error if there is no complete chain of nodes
leading to the node at the `idx`th index.
*/
func (l *linkedList[T]) DeleteAt(idx int) error {
	if idx == 0 {
		next := l.node.next
		if next != nil {
			l.node = next
		}

		return nil
	}

	n := l.node
	var previousNode *node[T]

	for range idx {
		previousNode = n
		n = n.next

		if n == nil {
			return fmt.Errorf("cannot delete node at index %d due to broken chain", idx)
		}
	}

	if n.next != nil {
		previousNode.next = n.next
	}

	return nil
}

/*
Reverse mutates the linkedList such that it's first `node` is now it's last
and it's last `node` now first. Values for each `node` are preserved.
*/
func (l *linkedList[T]) Reverse() {
	var prev *node[T]
	current := l.node

	for current != nil {
		next := current.next

		current.next = prev
		prev = current
		current = next
	}

	l.node = prev
}
