package datastructures

import "errors"

type queue struct {
	queue []int
}

func Queue() queue {
	return queue{
		queue: []int{},
	}
}

func (q *queue) Enqueue(num int) int {
	q.queue = append(q.queue, num)

	return num
}

func (q *queue) Dequeue() (int, error) {
	if len(q.queue) == 0 {
		return -1, errors.New("queue is empty")
	}

	first := q.queue[0]

	q.queue = q.queue[1:len(q.queue)]

	return first, nil
}

func (q *queue) Read() (int, error) {
	if len(q.queue) == 0 {
		return -1, errors.New("queue is empty")
	}

	return q.queue[0], nil
}
