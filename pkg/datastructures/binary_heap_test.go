package datastructures_test

import (
	"reflect"
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/datastructures"
)

func TestBinaryHeap_Insert(t *testing.T) {
	t.Parallel()

	t.Run("inserts lower value nodes", func(t *testing.T) {
		t.Parallel()

		rootNode := 5
		heap := datastructures.BinaryHeap(rootNode)

		for i := 1; i < 5; i++ {
			heap.Insert(i)
		}

		expectedHeap := []int{5, 4, 2, 1, 3}

		if heap.FirstNode() != rootNode {
			t.Fatalf("expected first node: %d, got: %d", rootNode, heap.FirstNode())
		}

		if !reflect.DeepEqual(heap.Show(), expectedHeap) {
			t.Fatalf("expected heap: %+v, got: %+v", expectedHeap, heap.Show())
		}

	})

	t.Run("replaces root node when inserted value is higher", func(t *testing.T) {
		t.Parallel()
		heap := datastructures.BinaryHeap(5)

		heap.Insert(8)

		if heap.FirstNode() != 8 {
			t.Fatalf("expected first node: 8, got: %d", heap.FirstNode())
		}

		heap.Insert(10)

		if heap.FirstNode() != 10 {
			t.Fatalf("expected first node: 10, got: %d", heap.FirstNode())
		}

		expectedHeap := []int{10, 5, 8}

		if !reflect.DeepEqual(heap.Show(), expectedHeap) {
			t.Fatalf("expected heap: %+v, got: %+v", expectedHeap, heap.Show())
		}

	})
}

func TestBinaryHeap_Delete(t *testing.T) {
	t.Parallel()

	t.Run("deletion reorders tree", func(t *testing.T) {
		t.Parallel()

		heap := datastructures.BinaryHeap(5)
		for i := 1; i < 5; i++ {
			heap.Insert(i)
		}

		expectedHeap := []int{5, 4, 2, 1, 3}

		if !reflect.DeepEqual(heap.Show(), expectedHeap) {
			t.Fatalf("expected deletion to result in: %+v, got: %+v", expectedHeap, heap.Show())
		}

		heap.Delete()

		expectedHeap = []int{4, 3, 2, 1}

		if !reflect.DeepEqual(heap.Show(), expectedHeap) {
			t.Fatalf("expected deletion to result in: %+v, got: %+v", expectedHeap, heap.Show())
		}
	})

	t.Run("deletion reorders on opposite side", func(t *testing.T) {
		t.Parallel()

		heap := datastructures.BinaryHeap(7)
		for i := 1; i < 7; i++ {
			heap.Insert(i)
		}

		expectedHeap := []int{7, 4, 6, 1, 3, 2, 5}

		if !reflect.DeepEqual(heap.Show(), expectedHeap) {
			t.Fatalf("expected deletion to result in: %+v, got: %+v", expectedHeap, heap.Show())
		}

		heap.Delete()

		expectedHeap = []int{6, 4, 5, 1, 3, 2}

		if !reflect.DeepEqual(heap.Show(), expectedHeap) {
			t.Fatalf("expected deletion to result in: %+v, got: %+v", expectedHeap, heap.Show())
		}
	})
}
