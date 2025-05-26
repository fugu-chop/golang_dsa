package datastructures_test

import (
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/datastructures"
)

func TestLinkedList_ReadandInsert(t *testing.T) {
	t.Parallel()

	t.Run("inserts when no element exists at next node", func(t *testing.T) {
		t.Parallel()

		linkedList := datastructures.LinkedList(0)

		if err := linkedList.InsertAt(1, 1); err != nil {
			t.Fatalf("error for InsertAt should be nil, got: %v", err)
		}

		data := map[string]struct {
			idx  int
			want int
		}{
			"index zero": {
				idx:  0,
				want: 0,
			},
			"index one": {
				idx:  1,
				want: 1,
			},
		}

		for _, testCase := range data {
			got, err := linkedList.ReadAt(testCase.idx)
			if err != nil {
				t.Fatalf("error for ReadAt should be nil, got: %v", err)
			}
			if got != testCase.want {
				t.Fatalf("value for ReadAt at index %d should be %d, got: %d", testCase.idx, testCase.want, got)
			}
		}
	})

	t.Run("inserts at index 0", func(t *testing.T) {
		t.Parallel()

		linkedList := datastructures.LinkedList(0)

		if err := linkedList.InsertAt(0, 1); err != nil {
			t.Fatalf("error for InsertAt should be nil, got: %v", err)
		}

		got, err := linkedList.ReadAt(0)
		if err != nil {
			t.Fatalf("error for ReadAt should be nil, got: %v", err)
		}
		if got != 1 {
			t.Fatalf("value for ReadAt should equal 1, got %d", got)
		}

		got, err = linkedList.ReadAt(1)
		if err == nil {
			t.Fatal("error for ReadAt should not be nil")
		}
		if got != -1 {
			t.Fatalf("returned index for ReadAt(1) should be -1, got %d", got)
		}
	})

	t.Run("inserts when next node already exists", func(t *testing.T) {
		t.Parallel()

		linkedList := datastructures.LinkedList(0)

		if err := linkedList.InsertAt(1, 1); err != nil {
			t.Fatalf("error for InsertAt should be nil, got: %v", err)
		}
		if err := linkedList.InsertAt(21, 2); err == nil {
			t.Fatal("error for InsertAt should not be nil")
		}
	})

	t.Run("returns an error when node chan is broken", func(t *testing.T) {
		t.Parallel()

		linkedList := datastructures.LinkedList(0)

		if err := linkedList.InsertAt(1, 1); err != nil {
			t.Fatalf("error for InsertAt should be nil, got: %v", err)
		}
		if err := linkedList.InsertAt(2, 2); err != nil {
			t.Fatalf("error for InsertAt should be nil, got: %v", err)
		}
		got, err := linkedList.ReadAt(1)
		if err != nil {
			t.Fatalf("error for ReadAt should be nil, got: %v", err)
		}
		if got != 1 {
			t.Fatalf("value for ReadAt at index 1 should be 1, got: %d", got)
		}

		if err := linkedList.InsertAt(1, 99); err != nil {
			t.Fatalf("error for InsertAt should be nil, got: %v", err)
		}
		got, err = linkedList.ReadAt(1)
		if err != nil {
			t.Fatalf("error for ReadAt should be nil, got: %v", err)
		}
		if got != 99 {
			t.Fatalf("value for ReadAt at index 1 should be 99, got: %d", got)
		}
		got, err = linkedList.ReadAt(2)
		if err != nil {
			t.Fatalf("error for ReadAt should be nil, got: %v", err)
		}
		if got != 1 {
			t.Fatalf("value for ReadAt at index 2 should be 1, got: %d", got)
		}
	})
}

func TestLinkedList_IndexOf(t *testing.T) {
	t.Parallel()

	t.Run("returns the index of a value that exists", func(t *testing.T) {
		t.Parallel()

		linkedList := datastructures.LinkedList(0)

		if err := linkedList.InsertAt(1, 97); err != nil {
			t.Fatalf("error for InsertAt should be nil, got: %v", err)
		}

		got, err := linkedList.IndexOf(97)
		if err != nil {
			t.Fatalf("error for IndexOf should be nil, got: %v", err)
		}
		if got != 1 {
			t.Fatalf("value for IndexOf should be 1, got: %d", got)
		}
	})

	t.Run("returns an error if the value does not exist", func(t *testing.T) {
		t.Parallel()

		linkedList := datastructures.LinkedList(0)

		if err := linkedList.InsertAt(1, 1); err != nil {
			t.Fatalf("error for InsertAt should be nil, got: %v", err)
		}

		got, err := linkedList.IndexOf(5)
		if err == nil {
			t.Fatal("error for IndexOf should be not be nil")
		}
		if got != -1 {
			t.Fatalf("value for IndexOf should be -1, got: %d", got)
		}
	})
}

func TestLinkedList_DeleteAt(t *testing.T) {
	t.Parallel()

	t.Run("deletes when element exists at node", func(t *testing.T) {
		t.Parallel()

		linkedList := datastructures.LinkedList(0)

		if err := linkedList.InsertAt(1, 1); err != nil {
			t.Fatalf("error for InsertAt should be nil, got: %v", err)
		}

		if err := linkedList.InsertAt(2, 2); err != nil {
			t.Fatalf("error for InsertAt should be nil, got: %v", err)
		}

		if err := linkedList.DeleteAt(1); err != nil {
			t.Fatalf("error for DeleteAt should be nil, got: %v", err)
		}

		got, err := linkedList.ReadAt(1)
		if err != nil {
			t.Fatalf("error for ReadAt should be nil, got: %v", err)
		}
		if got != 2 {
			t.Fatalf("ReadAt should have returned 2, got: %d", got)
		}

		got, err = linkedList.ReadAt(2)
		if err == nil {
			t.Fatalf("error for ReadAt should no be nil")
		}
		if got != -1 {
			t.Fatalf("ReadAt should have returned -1, got: %d", got)
		}
	})

	t.Run("returns an error when node does not exist", func(t *testing.T) {
		t.Parallel()

		linkedList := datastructures.LinkedList(0)

		if err := linkedList.DeleteAt(1); err == nil {
			t.Fatal("error for DeleteAt should not be nil")
		}

		if err := linkedList.DeleteAt(33); err == nil {
			t.Fatal("error for DeleteAt should not be nil")
		}
	})

	t.Run("handles zero index deletion", func(t *testing.T) {
		t.Parallel()

		linkedList := datastructures.LinkedList(0)

		if err := linkedList.InsertAt(1, 1); err != nil {
			t.Fatalf("error for InsertAt should be nil, got: %v", err)
		}

		if err := linkedList.DeleteAt(0); err != nil {
			t.Fatalf("error for DeleteAt should be nil, got: %v", err)
		}

		got, err := linkedList.ReadAt(0)
		if err != nil {
			t.Fatalf("error for ReadAt should be nil, got: %v", err)
		}
		if got != 1 {
			t.Fatalf("ReadAt should be 1, got: %d", got)
		}
	})
}
