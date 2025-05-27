package datastructures_test

import (
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/datastructures"
)

func TestStackPush(t *testing.T) {
	t.Parallel()

	stack := datastructures.Stack(0)
	for i := 1; i < 10; i++ {
		got := stack.Push(i)

		if got != i {
			t.Fatalf("stack did not Push correctly: got: %d, want: %d", got, i)
		}

		read, err := stack.Read()
		if read != i {
			t.Fatalf("stack did not read correctly: got: %d, want: %d", got, i)
		}
		if err != nil {
			t.Fatalf("should not have gottten error from read: got: %v", err)
		}
	}
}

func TestStackPop(t *testing.T) {
	t.Parallel()

	t.Run("populated stack", func(t *testing.T) {
		t.Parallel()

		stack := datastructures.Stack(0)
		for i := 1; i <= 4; i++ {
			_ = stack.Push(i)
		}

		for i := 4; i > 0; i-- {
			got, err := stack.Pop()
			if got != i {
				t.Fatalf(
					"Pop on stack did not return result in order, expected: %d, got: %d", i, got,
				)
			}
			if err != nil {
				t.Fatalf("should not have gottten error from read: got: %v", err)
			}

			/*
			 Read should now have the updated value - terminate early since there will be
			 one fewer Read available than Pop.
			*/
			read, err := stack.Read()
			if err != nil {
				t.Fatalf("should not have gottten error from read: got: %v", err)
			}
			if read != got-1 {
				t.Fatalf("Pop on stack did not update last node - got: %d, want: %d", read, got-1)
			}
		}
	})

	t.Run("empty stack", func(t *testing.T) {
		t.Parallel()

		stack := datastructures.Stack(0)
		_, _ = stack.Pop()

		got, err := stack.Pop()
		if got != -1 {
			t.Fatalf("Pop should return -1 on empty stack, got: %d", got)
		}
		if err == nil {
			t.Fatal("should have gottten error from Pop on empty stack")
		}
	})
}

func TestStackRead(t *testing.T) {
	t.Parallel()

	t.Run("populated stack", func(t *testing.T) {
		t.Parallel()

		stack := datastructures.Stack(0)

		for i := 1; i < 5; i++ {
			_ = stack.Push(i)
		}

		for range 10 {
			got, err := stack.Read()
			// Read is non-destructive and idempotent
			if got != 4 {
				t.Fatalf("Read did not return correct result - got: %d, want: %d", got, 4)
			}
			if err != nil {
				t.Fatalf("should not have gottten error from read: got: %v", err)
			}
		}
	})

	t.Run("empty stack", func(t *testing.T) {
		t.Parallel()

		stack := datastructures.Stack(0)

		_, _ = stack.Pop()
		got, err := stack.Read()
		// Read is non-destructive and idempotent
		if err == nil {
			t.Fatal("should have gottten error from Read on empty stack")
		}
		if got != -1 {
			t.Fatalf("Read did not return correct result - got: %d, want: %d", got, 4)
		}
	})
}
