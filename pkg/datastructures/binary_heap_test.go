package datastructures

import (
	"reflect"
	"testing"
)

/*
show returns the entire heap. It is used primarily for debugging and testing purposes.
*/
func show(t *testing.T, b *binaryHeap) []int {
	t.Helper()

	return b.heap
}

func TestBinaryHeap_Insert(t *testing.T) {
	t.Parallel()

	t.Run("inserts lower value nodes", func(t *testing.T) {
		t.Parallel()

		rootNode := 5
		heap := BinaryHeap(rootNode)

		for i := 1; i < 5; i++ {
			heap.Insert(i)
		}

		expectedHeap := []int{5, 4, 2, 1, 3}

		if heap.FirstNode() != rootNode {
			t.Fatalf("expected first node: %d, got: %d", rootNode, heap.FirstNode())
		}

		if !reflect.DeepEqual(show(t, heap), expectedHeap) {
			t.Fatalf("expected heap: %+v, got: %+v", expectedHeap, show(t, heap))
		}

	})

	t.Run("replaces root node when inserted value is higher", func(t *testing.T) {
		t.Parallel()
		heap := BinaryHeap(5)

		heap.Insert(8)

		if heap.FirstNode() != 8 {
			t.Fatalf("expected first node: 8, got: %d", heap.FirstNode())
		}

		heap.Insert(10)

		if heap.FirstNode() != 10 {
			t.Fatalf("expected first node: 10, got: %d", heap.FirstNode())
		}

		expectedHeap := []int{10, 5, 8}

		if !reflect.DeepEqual(show(t, heap), expectedHeap) {
			t.Fatalf("expected heap: %+v, got: %+v", expectedHeap, show(t, heap))
		}

	})
}

func TestBinaryHeap_Delete(t *testing.T) {
	t.Parallel()

	t.Run("deletion reorders tree", func(t *testing.T) {
		t.Parallel()

		rootVal := 5

		heap := BinaryHeap(rootVal)
		for i := 1; i < 5; i++ {
			heap.Insert(i)
		}

		expectedHeap := []int{5, 4, 2, 1, 3}

		if !reflect.DeepEqual(show(t, heap), expectedHeap) {
			t.Fatalf("expected deletion to result in: %+v, got: %+v", expectedHeap, show(t, heap))
		}

		got := heap.Delete()
		if *got != rootVal {
			t.Fatalf("expected delete to return root node value of %d, got: %d", rootVal, got)
		}

		expectedHeap = []int{4, 3, 2, 1}

		if !reflect.DeepEqual(show(t, heap), expectedHeap) {
			t.Fatalf("expected deletion to result in: %+v, got: %+v", expectedHeap, show(t, heap))
		}
	})

	t.Run("deletion reorders on opposite side", func(t *testing.T) {
		t.Parallel()

		rootVal := 7

		heap := BinaryHeap(rootVal)
		for i := 1; i < 7; i++ {
			heap.Insert(i)
		}

		expectedHeap := []int{7, 4, 6, 1, 3, 2, 5}

		if !reflect.DeepEqual(show(t, heap), expectedHeap) {
			t.Fatalf("expected deletion to result in: %+v, got: %+v", expectedHeap, show(t, heap))
		}

		got := heap.Delete()
		if *got != rootVal {
			t.Fatalf("expected delete to return root node value of %d, got: %d", rootVal, got)
		}

		expectedHeap = []int{6, 4, 5, 1, 3, 2}

		if !reflect.DeepEqual(show(t, heap), expectedHeap) {
			t.Fatalf("expected deletion to result in: %+v, got: %+v", expectedHeap, show(t, heap))
		}
	})

	t.Run("handles no nodes", func(t *testing.T) {
		t.Parallel()

		rootVal := 7
		heap := BinaryHeap(rootVal)

		got := heap.Delete()
		if *got != rootVal {
			t.Fatalf("expected delete to return root node value of %d, got: %d", rootVal, got)
		}

		expectedHeap := []int{}

		if !reflect.DeepEqual(show(t, heap), expectedHeap) {
			t.Fatalf("expected deletion to result in: %+v, got: %+v", expectedHeap, show(t, heap))
		}

		got = heap.Delete()
		if got != nil {
			t.Fatalf("expected delete to return root node value of nil, got: %d", *got)
		}

		expectedHeap = []int{}

		if !reflect.DeepEqual(show(t, heap), expectedHeap) {
			t.Fatalf("expected deletion to result in: %+v, got: %+v", expectedHeap, show(t, heap))
		}
	})
}
