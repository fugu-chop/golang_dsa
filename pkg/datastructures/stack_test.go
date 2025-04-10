package datastructures_test

import (
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/datastructures"
)

func TestStackPush(t *testing.T) {
	t.Parallel()

	stack := datastructures.Stack()
	for i := 0; i < 10; i++ {
		got := stack.Push(i)

		if got != i {
			t.Fatalf("Stack did not push correctly: got: %d, want: %d", got, i)
		}
	}
}

func TestStackPop(t *testing.T) {
	t.Parallel()

	t.Run("empty stack", func(t *testing.T) {
		t.Parallel()

		stack := datastructures.Stack()
		result, err := stack.Pop()

		if err == nil {
			t.Fatal("Pop on empty stack should have returned an error")
		}
		if err.Error() != "stack is empty" {
			t.Fatalf("Pop on empty stack should have returned \"stack is empty\" error message, got: %s", err.Error())
		}
		if result != -1 {
			t.Fatalf("Pop on empty stack did not return -1, returned %d instead", result)
		}
	})

	t.Run("populated stack", func(t *testing.T) {
		t.Parallel()

		stack := datastructures.Stack()

		for i := 0; i < 10; i++ {
			_ = stack.Push(i)
		}

		for i := 9; i > -1; i-- {
			got, err := stack.Pop()

			if got != i {
				t.Fatalf(
					"Pop on non-empty stack did not return result in order, expected: %d, got: %d",
					i,
					got,
				)
			}

			if err != nil {
				t.Fatalf("Pop on non-empty stack returned error: %+v", err)
			}
		}
	})
}

func TestStackRead(t *testing.T) {
	t.Parallel()

	t.Run("empty read", func(t *testing.T) {
		t.Parallel()

		stack := datastructures.Stack()
		result, err := stack.Read()

		if result != -1 {
			t.Fatalf("Read on empty stack did not return -1, returned: %d", result)
		}

		if err == nil {
			t.Fatal("Read on empty stack did not return error")
		}

		if err.Error() != "stack is empty" {
			t.Fatalf("Read on empty stack did not return correct error message: %s", err.Error())
		}
	})

	t.Run("populated stack", func(t *testing.T) {
		t.Parallel()

		stack := datastructures.Stack()

		for i := 0; i < 5; i++ {
			_ = stack.Push(i)
		}

		for i := 5; i > -1; i-- {
			got, err := stack.Read()
			if err != nil {
				t.Fatalf("Read failed: %v", err)
			}

			// Read is non-destructive and idempotent
			if got != 4 {
				t.Fatalf("Read did not return correct result - got: %d, want: %d", got, 4)
			}
		}
	})
}
