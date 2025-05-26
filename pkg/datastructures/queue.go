package datastructures

type queue struct {
	firstNode *node
	lastNode  *node
}

func Queue(val int) queue {
	newNode := &node{
		value: &val,
	}

	return queue{
		firstNode: newNode,
		lastNode:  newNode,
	}
}

func (q *queue) Enqueue(num int) int {
	newNode := &node{
		value: &num,
	}

	q.lastNode.next = newNode
	q.lastNode = newNode

	return num
}

func (q *queue) Dequeue() (int, error) {
	val := *q.firstNode.value
	next := q.firstNode.next
	q.firstNode = next

	penultimateNode := q.firstNode
	for penultimateNode.next != nil {
		penultimateNode = penultimateNode.next
	}

	q.lastNode = penultimateNode

	return val, nil
}

func (q *queue) Read() (int, error) {
	return *q.firstNode.value, nil
}
