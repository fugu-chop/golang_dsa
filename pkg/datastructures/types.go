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
It contains a slice of integers instead of references to nodes to
allow easier access to the last 'node' in the heap.
This makes deletion of nodes much easier.
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

/*
trie is an implementation of a trie (retrieval) data structure.
It holds a `root`, which is a pointer to a trieNode.
*/
type trie struct {
	root *trieNode
}

/*
a trieNode is a node used within the implementation of a Trie.
It contains a map whose keys are letters and values are pointers to
other trieNodes.
*/
type trieNode struct {
	children map[string]*trieNode
}

/*
Get attempts to fetch a pointer to a child node that has a child with a value of `letter`.
*/
func (t *trieNode) get(letter string) *trieNode {
	return t.children[letter]
}

/*
Set creates a child node for the current node that has a `letter` value.
If a child node already exists with a `letter` value, `Set` is a no-op.
*/
func (t *trieNode) set(letter string) {
	// avoid clobbering existing relationships
	if _, ok := t.children[letter]; ok {
		return
	}

	t.children[letter] = &trieNode{
		children: make(map[string]*trieNode),
	}
}
