package datastructures

import (
	"fmt"
)

// NewLinkedList returns a pointer to a new linkedList.
func LinkedList(val int) *linkedList {
	return &linkedList{
		node: &node{
			value: val,
		},
	}
}

/*
ReadAt follows the list of nodes in the Linked List until it
reaches the node at index `idx`.

It returns an error if there is no node at index `idx`.
*/
func (l *linkedList) ReadAt(idx int) (int, error) {
	n := l.node
	for range idx {
		if n.next == nil {
			return -1, fmt.Errorf("node at index %d does not exist", idx)
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
func (l *linkedList) IndexOf(value int) (int, error) {
	idx := 0
	n := l.node

	for n != nil {
		if n.value == value {
			return idx, nil
		}
		n = n.next
		idx++
	}

	return -1, fmt.Errorf("value '%d' does not exist in linked list", value)
}

/*
InsertAt inserts a node containing a value of `value` at the `idx`th
node of the Linked List.

It returns an error if there is no complete chain of nodes leading to
the node at index `idx`.
*/
func (l *linkedList) InsertAt(idx int, val int) error {
	newNode := &node{
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
func (l *linkedList) DeleteAt(idx int) error {
	if idx == 0 {
		next := l.node.next
		if next != nil {
			l.node = next
		}

		return nil
	}

	n := l.node
	var previousNode *node

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
func (l *linkedList) Reverse() {
	var prev *node
	current := l.node

	for current != nil {
		next := current.next

		current.next = prev
		prev = current
		current = next
	}

	l.node = prev
}
