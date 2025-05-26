package datastructures

/*
The node type is intended to be used as the node component of the
LinkedList, Queue and Stack types. Each node has two attributes:

1. An int; and

2. A pointer to the next `node`.
*/
type node struct {
	value int
	next  *node
}
