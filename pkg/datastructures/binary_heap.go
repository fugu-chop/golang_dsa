package datastructures

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
BinaryHeap returns a pointer to a binaryHeap type
with the `heap` populated with a single value, `val`.
*/
func BinaryHeap(val int) *binaryHeap {
	return &binaryHeap{
		heap: []int{val},
	}
}

/*
FirstNode returns the first element in the heap (i.e. the root node).
*/
func (b *binaryHeap) FirstNode() int {
	return b.heap[0]
}

/*
Insert adds a node to the heap and trickles it upwards based on the
value of parent nodes - i.e. Insert works on a max heap basis where
the largest value in the entire heap should be the root node.
*/
func (b *binaryHeap) Insert(val int) {
	b.heap = append(b.heap, val)
	newNodeIdx := len(b.heap) - 1

	// OOB guard as root node will not have a parent
	for newNodeIdx > 0 && b.heap[newNodeIdx] > b.heap[parentIdx(newNodeIdx)] {
		b.heap[newNodeIdx], b.heap[parentIdx(newNodeIdx)] =
			b.heap[parentIdx(newNodeIdx)], b.heap[newNodeIdx]
		newNodeIdx = parentIdx(newNodeIdx)
	}
}

/*
Delete removes the root node and replaces it with the next highest
value within the heap (i.e. a max heap). It does this by replacing the
root node with the lastNode and then trickling it downward based on the
value of it's child nodes.

It returns a pointer to the value of the original root node (before deletion).
*/
func (b *binaryHeap) Delete() *int {
	if len(b.heap) == 0 {
		return nil
	}

	rootVal := b.FirstNode()
	b.heap[0] = b.lastNode()
	b.heap = b.heap[:len(b.heap)-1]
	trickleNodeIdx := 0

	for b.hasGreaterChild(trickleNodeIdx) {
		largerChildIdx := b.largerChildNodeIdx(trickleNodeIdx)
		b.heap[trickleNodeIdx], b.heap[largerChildIdx] =
			b.heap[largerChildIdx], b.heap[trickleNodeIdx]
		trickleNodeIdx = largerChildIdx
	}

	return &rootVal
}

/*
Show returns the entire heap. It is used primarily for debugging and testing purposes.
*/
func (b *binaryHeap) Show() []int {
	return b.heap
}

/*
lastNode returns the last element in the heap (i.e. the last element in the slice).
*/
func (b *binaryHeap) lastNode() int {
	return b.heap[len(b.heap)-1]
}

/*
hasGreaterChild determines whether or not a given node has a child node with a value
greater than it's own value. It has a built in out of bounds guard to avoid panics
where no child exists for a given node.
*/
func (b *binaryHeap) hasGreaterChild(idx int) bool {
	return leftChildIdx(idx) < len(b.heap) &&
		b.heap[leftChildIdx(idx)] > b.heap[idx] ||
		rightChildIdx(idx) < len(b.heap) &&
			b.heap[rightChildIdx(idx)] > b.heap[idx]
}

/*
largerChildNodeIdx returns the index of a node where that node's value is
larger than it's parent.
*/
func (b *binaryHeap) largerChildNodeIdx(idx int) int {
	// In theory this should be an out of bounds check
	// but I cannot get a test case to simulate this

	// No right child
	if rightChildIdx(idx) >= len(b.heap) {
		return leftChildIdx(idx)
	}

	// existence of a right child implies a left child due to
	// completeness condition of binary heap
	if b.heap[rightChildIdx(idx)] > b.heap[leftChildIdx(idx)] {
		return rightChildIdx(idx)
	}

	return leftChildIdx(idx)
}

/*
parentIdx returns the index of a particular node's parent.
It relies on integer division (i.e. decimal places are ignored)
and operates on the assumption of a maximum of two child nodes per node
(i.e. a 2^n of nodes, where n == number of levels).
*/
func parentIdx(idx int) int {
	return (idx - 1) / 2
}

/*
leftChildIdx returns the index of a particular node's 'left child'.
It relies on integer division (i.e. decimal places are ignored)
and operates on the assumption of a maximum of two child nodes per node
(i.e. a 2^n of nodes, where n == number of levels).
*/
func leftChildIdx(idx int) int {
	return 2*idx + 1
}

/*
rightChildIdx returns the index of a particular node's 'right child'.
It relies on integer division (i.e. decimal places are ignored)
and operates on the assumption of a maximum of two child nodes per node
(i.e. a 2^n of nodes, where n == number of levels).
*/
func rightChildIdx(idx int) int {
	return 2*idx + 2
}
