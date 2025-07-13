package datastructures

import "errors"

/*
queue is an implementation of a queue data structure. It contains a pointer
to the first `node` type and the last `node` type it encompasses.
*/
type queue[T any] struct {
	firstNode *node[T]
	lastNode  *node[T]
}

// Queue returns a pointer to a new queue with a `node` of `val` value.
func Queue[T any](val T) queue[T] {
	newNode := &node[T]{
		value: val,
	}

	return queue[T]{
		firstNode: newNode,
		lastNode:  newNode,
	}
}

/*
Enqueue adds a pointer to a new `node` (with the value of `num`) to the end of the queue.
Enqueue replaces the pointer of the `lastNode` to the newly created `node`.
*/
func (q *queue[T]) Enqueue(val T) T {
	newNode := &node[T]{
		value: val,
	}

	q.lastNode.next = newNode
	q.lastNode = newNode

	return val
}

/*
Dequeue removes the pointer to the current `firstNode` and replaces it with
a pointer to the node that was it's previous `next` node. It returns the value
of the current `firstNode` (before pointer reassignment), or an error if there
are no nodes in the queue.
*/
func (q *queue[T]) Dequeue() (T, error) {
	if q.firstNode == nil {
		return *new(T), errors.New("queue does not have any active nodes")
	}

	val := q.firstNode.value
	next := q.firstNode.next
	q.firstNode = next

	return val, nil
}

/*
Read returns the value associated with the `node` at the `firstNode` reference,
or an error if there are no nodes in the queue.
*/
func (q *queue[T]) Read() (T, error) {
	if q.firstNode == nil {
		return *new(T), errors.New("queue does not have any active nodes")
	}

	return q.firstNode.value, nil
}
