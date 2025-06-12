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
any of the linkedList, stack or queue, but primarily intended
for usage in a binarySearchTree. Each node has the same
attributes as `node` but contains an additional pointer to
the previous `doubleLinkedNode` in the chain.
*/
type doubleLinkedNode struct {
	Value int
	next  *doubleLinkedNode
	prev  *doubleLinkedNode
}

/*
binaryHeap is an implementation of a Binary Heap data structure.
It contains a slice of integers.
*/
type binaryHeap struct {
	heap []int
}

/*
binarySearchTree is an implementation of a Binary Search Tree data structure.
It contains a pointer to a doubleLinkedNode type.
*/
type binarySearchTree struct {
	CurrentNode *doubleLinkedNode
}

/*
linkedList is an implementation of a single Linked List
data structure. It contains a pointer to a `node` type.
*/
type linkedList struct {
	node *node
}

/*
queue is an implementation of a queue data structure. It contains a pointer
to the first `node` type and the last `node` type it encompasses.
*/
type queue struct {
	firstNode *node
	lastNode  *node
}

/*
stack is an implementation of a stack data structure. It contains a pointer
to a `node`.
*/
type stack struct {
	currentNode *node
}
