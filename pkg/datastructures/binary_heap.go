package datastructures

type binaryHeap struct {
	heap []int
}

func BinaryHeap(val int) *binaryHeap {
	return &binaryHeap{
		heap: []int{val},
	}
}

func (b *binaryHeap) FirstNode() int {
	return b.heap[0]
}

func (b *binaryHeap) Insert(val int) {
	b.heap = append(b.heap, val)
	newNodeIdx := len(b.heap) - 1

	// OOB guard as root node will not have a parent
	for newNodeIdx > 0 &&
		b.heap[newNodeIdx] < b.heap[parentIdx(newNodeIdx)] {
		b.heap[newNodeIdx], b.heap[parentIdx(newNodeIdx)] =
			b.heap[parentIdx(newNodeIdx)], b.heap[newNodeIdx]
		newNodeIdx = parentIdx(newNodeIdx)
	}
}

func (b *binaryHeap) Delete() {
	b.heap[0] = b.lastNode()
	b.heap = b.heap[:len(b.heap)-1]
	trickleNodeIdx := 0

	for b.hasGreaterChild(trickleNodeIdx) {
		largerChildIdx := b.largerChildNodeIdx(trickleNodeIdx)
		b.heap[trickleNodeIdx], b.heap[largerChildIdx] =
			b.heap[largerChildIdx], b.heap[trickleNodeIdx]
		trickleNodeIdx = largerChildIdx
	}
}

func (b *binaryHeap) lastNode() int {
	return b.heap[len(b.heap)]
}

func (b *binaryHeap) hasGreaterChild(idx int) bool {
	return leftChildIdx(idx) < len(b.heap) &&
		b.heap[leftChildIdx(idx)] > b.heap[idx] ||
		rightChildIdx(idx) < len(b.heap) &&
			b.heap[rightChildIdx(idx)] > b.heap[idx]
}

func (b *binaryHeap) largerChildNodeIdx(idx int) int {
	// No right child
	if rightChildIdx(idx) >= len(b.heap) {
		return leftChildIdx(idx)
	}

	if b.heap[rightChildIdx(idx)] > b.heap[leftChildIdx(idx)] {
		return rightChildIdx(idx)
	}

	return leftChildIdx(idx)
}

// Relies on integer division (i.e. decimal places are ignored).
func parentIdx(idx int) int {
	return (idx - 1) / 2
}

func leftChildIdx(idx int) int {
	return 2*idx + 1
}

func rightChildIdx(idx int) int {
	return 2*idx + 2
}
