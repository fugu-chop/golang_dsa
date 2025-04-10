package datastructures_test

import (
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/datastructures"
)

func TestQueueEnqueue(t *testing.T) {
	t.Parallel()

	queue := datastructures.Queue()

	for i := 0; i < 5; i++ {
		got := queue.Enqueue(i)
		if got != i {
			t.Fatalf("Queue did not Enqueue correctly: got: %d, want: %d", got, i)
		}
	}
}

func TestQueueDequeue(t *testing.T) {
	t.Parallel()

	t.Run("empty queue", func(t *testing.T) {
		t.Parallel()

		queue := datastructures.Queue()
		result, err := queue.Dequeue()

		if err == nil {
			t.Fatal("Dequeue on empty queue should have returned an error")
		}
		if err.Error() != "queue is empty" {
			t.Fatalf("Dequeue on empty queue should have returned \"queue is empty\" error message, got: %s", err.Error())
		}
		if result != -1 {
			t.Fatalf("Dequeue on empty queue did not return -1, returned %d instead", result)
		}
	})

	t.Run("populated queue", func(t *testing.T) {
		t.Parallel()

		queue := datastructures.Queue()

		for i := 0; i < 10; i++ {
			_ = queue.Enqueue(i)
		}

		for i := 0; i < 10; i++ {
			got, err := queue.Dequeue()

			if got != i {
				t.Fatalf("Dequeue on non-empty queue did not return result in order, expected: %d, got: %d", i, got)
			}

			if err != nil {
				t.Fatalf("Dequeue on non-empty queue returned error: %+v", err)
			}
		}
	})
}

func TestQueueRead(t *testing.T) {
	t.Parallel()

	t.Run("empty queue", func(t *testing.T) {
		t.Parallel()
		queue := datastructures.Queue()
		result, err := queue.Read()

		if result != -1 {
			t.Fatalf("Read on empty queue did not return -1, returned: %d", result)
		}

		if err == nil {
			t.Fatal("Read on empty queue did not return error")
		}

		if err.Error() != "queue is empty" {
			t.Fatalf("Read on empty queue did not return correct error message: %s", err.Error())
		}
	})

	t.Run("populated queue", func(t *testing.T) {
		t.Parallel()

		queue := datastructures.Queue()

		for i := 0; i < 10; i++ {
			_ = queue.Enqueue(i)
		}

		for i := 0; i < 5; i++ {
			result, err := queue.Read()
			if err != nil {
				t.Fatalf("Read returned an error: %v", err)
			}
			if result != 0 {
				t.Fatalf("Read should not mutate the queue - got: %d, want 0", result)
			}
		}
	})
}
