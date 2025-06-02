package datastructures_test

import (
	"math/rand"
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
