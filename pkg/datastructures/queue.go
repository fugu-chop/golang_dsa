package datastructures

/*
queue is an implementation of a queue data structure. It contains a pointer
to the first `node` type and the last `node` type it encompasses.
*/
type queue struct {
	firstNode *node
	lastNode  *node
}

// Queue returns a pointer to a new queue.
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
of the current `firstNode` (before pointer reassignment).
*/
func (q *queue) Dequeue() int {
	val := q.firstNode.value
	next := q.firstNode.next
	q.firstNode = next

	return val
}

// Read returns the value associated with the `node` at the `firstNode` reference.
func (q *queue) Read() int {
	return q.firstNode.value
}
