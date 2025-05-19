package datastructures

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

func (l *LinkedList) ReadAt(idx int) *int {
	n := l.node
	for range idx {
		if n.next == nil {
			return nil
		}
		n = n.next
	}

	return n.value
}

func (l *LinkedList) Search(value int) int {
	idx := 0
	n := l.node

	for n.next != nil {
		if *n.value == value {
			return idx
		}
		n = n.next
		idx++
	}

	return -1
}
