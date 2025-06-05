package datastructures_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/datastructures"
)

func TestBinarySearchTree_InsertAndSearch(t *testing.T) {
	t.Parallel()

	t.Run("finds nodes when they exist", func(t *testing.T) {
		t.Parallel()

		randInts := []int{}
		for range 20 {
			randInts = append(randInts, rand.Intn(100))
		}

		b := datastructures.BinarySearchTree(randInts[0])

		for _, num := range randInts[1:] {
			b.Insert(num, b.CurrentNode)
		}

		for _, num := range randInts {
			result := b.Search(num, b.CurrentNode)
			if result == nil {
				t.Fatalf("expected to find %v with value %d in tree, got nil", result, num)
			}
		}
	})

	t.Run("returns nil when no node exists", func(t *testing.T) {
		t.Parallel()

		b := datastructures.BinarySearchTree(5)

		result := b.Search(3, b.CurrentNode)
		if result != nil {
			t.Fatalf("found node when it should not exist for value: %d", 3)
		}
	})
}

func TestBinarySearchTree_Delete(t *testing.T) {
	t.Parallel()

	t.Run("returns root node when node does not exist", func(t *testing.T) {
		t.Parallel()

		b := datastructures.BinarySearchTree(3)

		result := b.Delete(999, b.CurrentNode)
		if result != b.CurrentNode {
			t.Fatalf("Delete on non-existing node should return nil, got: %v", result)
		}
		if b.CurrentNode.Value != 3 {
			t.Fatalf("expected current node to be %d, got: %d", 3, b.CurrentNode.Value)
		}
	})

	t.Run("handles deletion of nodes with no children", func(t *testing.T) {
		t.Parallel()

		vals := []int{1, 6, 5, 8, 2, 3}
		b := datastructures.BinarySearchTree(4)
		for _, val := range vals {
			b.Insert(val, b.CurrentNode)
		}

		node := b.Delete(8, b.CurrentNode)

		if node != b.CurrentNode {
			t.Fatalf("Delete on non-existing node should return nil, got: %v", node)
		}

		traversal := b.Traverse(b.CurrentNode, []int{})
		expectedTraversal := []int{4, 1, 2, 3, 6, 5}
		if !reflect.DeepEqual(expectedTraversal, traversal) {
			t.Fatalf("deletion of node should prune tree, want: %v, got: %v", expectedTraversal, traversal)
		}
	})

	t.Run("handles deletion of nodes with one left child", func(t *testing.T) {
		t.Parallel()

		vals := []int{1, 6, 5, 8, 2}
		b := datastructures.BinarySearchTree(4)
		for _, val := range vals {
			b.Insert(val, b.CurrentNode)
		}

		node := b.Delete(1, b.CurrentNode)

		if node != b.CurrentNode {
			t.Fatalf("Delete on non-existing node should return nil, got: %v", node)
		}

		traversal := b.Traverse(b.CurrentNode, []int{})
		expectedTraversal := []int{4, 2, 6, 5, 8}
		if !reflect.DeepEqual(expectedTraversal, traversal) {
			t.Fatalf("deletion of node should prune tree, want: %v, got: %v", expectedTraversal, traversal)
		}
	})

	t.Run("handles deletion of nodes with one right child", func(t *testing.T) {
		t.Parallel()

		vals := []int{1, 6, 5, 2, 3}
		b := datastructures.BinarySearchTree(4)
		for _, val := range vals {
			b.Insert(val, b.CurrentNode)
		}

		node := b.Delete(6, b.CurrentNode)

		if node != b.CurrentNode {
			t.Fatalf("Delete on non-existing node should return nil, got: %v", node)
		}

		traversal := b.Traverse(b.CurrentNode, []int{})
		expectedTraversal := []int{4, 1, 2, 3, 5}
		if !reflect.DeepEqual(expectedTraversal, traversal) {
			t.Fatalf("deletion of node should prune tree, want: %v, got: %v", expectedTraversal, traversal)
		}
	})

	t.Run("handles deletion of root node (two children)", func(t *testing.T) {
		t.Parallel()

		rootNodeVal := 4
		vals := []int{1, 6, 5, 8, 2, 3}
		b := datastructures.BinarySearchTree(rootNodeVal)
		for _, val := range vals {
			b.Insert(val, b.CurrentNode)
		}

		node := b.Delete(rootNodeVal, b.CurrentNode)
		if node.Value == rootNodeVal {
			t.Fatalf("root node should have been deleted")
		}
		if b.CurrentNode.Value == rootNodeVal {
			t.Fatalf("currentNode value should have been replaced with %d, got: %d", rootNodeVal, b.CurrentNode.Value)
		}
		traversal := b.Traverse(b.CurrentNode, []int{})
		expectedTraversal := []int{5, 1, 2, 3, 6, 8}
		if !reflect.DeepEqual(expectedTraversal, traversal) {
			t.Fatalf("deletion of node should prune tree, want: %v, got: %v", expectedTraversal, traversal)
		}
	})
}

func TestBinarySearchTree_Traverse(t *testing.T) {
	t.Parallel()

	vals := []int{1, 6, 5, 8, 2, 3}
	b := datastructures.BinarySearchTree(4)
	for _, val := range vals {
		b.Insert(val, b.CurrentNode)
	}

	result := b.Traverse(b.CurrentNode, []int{})
	expected := []int{4, 1, 2, 3, 6, 5, 8}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected result to equal %v, got: %v", expected, result)
	}
}
