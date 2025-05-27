package datastructures

/*
The `node` type is intended to be used as the node component of the
linkedList, stack and queue types. Each node has two attributes:

1. An int; and

2. A pointer to the next `node`.
*/
type node struct {
	value int
	next  *node
}

/*
The `doubleLinkedNode` can be used as the node component of
any of the linkedList, stack or queue. Each node has the same
attributes as `node` but contains an additional pointer to
the previous `doubleLinkedNode` in the chain.
*/
type doubleLinkedNode struct {
	value int
	next  *doubleLinkedNode
	prev  *doubleLinkedNode
}
