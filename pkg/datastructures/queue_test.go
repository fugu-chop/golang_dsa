package datastructures_test

import (
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/datastructures"
)

func TestQueueEnqueue(t *testing.T) {
	t.Parallel()

	queue := datastructures.Queue(0)

	for i := 1; i < 5; i++ {
		got := queue.Enqueue(i)
		if got != i {
			t.Fatalf("Queue did not Enqueue correctly: got: %d, want: %d", got, i)
		}
	}
}

func TestQueueDequeue(t *testing.T) {
	t.Parallel()

	t.Run("populated queue", func(t *testing.T) {
		t.Parallel()

		queue := datastructures.Queue(0)

		for i := 1; i < 10; i++ {
			_ = queue.Enqueue(i)
		}

		for i := range 9 {
			got, err := queue.Dequeue()

			if got != i {
				t.Fatalf("Dequeue on non-empty queue did not return result in order, expected: %d, got: %d", i, got)
			}
			if err != nil {
				t.Fatalf("should not have errors on Dequeue, got: %v", err)
			}

			// Read should now read off the new firstNode
			got, err = queue.Read()
			if i+1 != got {
				t.Fatalf("last node was not correctly replaced, got: %d, want: %d", got, i+1)
			}
			if err != nil {
				t.Fatalf("should not have errors on Read, got: %v", err)
			}
		}
	})

	t.Run("empty queue", func(t *testing.T) {
		queue := datastructures.Queue(0)

		_, _ = queue.Dequeue()

		for range 100 {
			got, err := queue.Dequeue()

			if got != -1 {
				t.Fatalf("Dequeue on empty queue should return -1, got: %d", got)
			}
			if err == nil {
				t.Fatalf("should have errors on Dequeue of empty queue")
			}
		}
	})
}

func TestQueueRead(t *testing.T) {
	t.Parallel()

	queue := datastructures.Queue(0)

	for i := range 5 {
		_ = queue.Enqueue(i)
	}

	for range 100 {
		result, err := queue.Read()
		if result != 0 {
			t.Fatalf("Read should not mutate the queue - got: %d, want 0", result)
		}
		if err != nil {
			t.Fatalf("should not have errors on Read, got: %v", err)
		}
	}
}
