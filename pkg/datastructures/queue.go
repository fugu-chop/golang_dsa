package datastructures

import "errors"

// Queue returns a pointer to a new queue with a `node` of `val` value.
func Queue(val int) queue {
	newNode := &node{
		value: val,
	}

	return queue{
		firstNode: newNode,
		lastNode:  newNode,
	}
}

/*
Enqueue adds a pointer to a new `node` (with the value of `num`) to the end of the queue.
Enqueue replaces the pointer of the `lastNode` to the newly created `node`.
*/
func (q *queue) Enqueue(num int) int {
	newNode := &node{
		value: num,
	}

	q.lastNode.next = newNode
	q.lastNode = newNode

	return num
}

/*
Dequeue removes the pointer to the current `firstNode` and replaces it with
a pointer to the node that was it's previous `next` node. It returns the value
of the current `firstNode` (before pointer reassignment), or an error if there
are no nodes in the queue.
*/
func (q *queue) Dequeue() (int, error) {
	if q.firstNode == nil {
		return -1, errors.New("queue does not have any active nodes")
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
func (q *queue) Read() (int, error) {
	if q.firstNode == nil {
		return -1, errors.New("queue does not have any active nodes")
	}

	return q.firstNode.value, nil
}
