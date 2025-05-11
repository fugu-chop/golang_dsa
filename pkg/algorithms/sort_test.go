package algorithms_test

import (
	"reflect"
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/algorithms"
)

func TestBubbleSort(t *testing.T) {
	t.Parallel()

	data := map[string]struct {
		inputSlice  []int
		outputSlice []int
	}{
		"sorts": {
			inputSlice:  []int{9, 7, 5, 6, 1, 2},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"sorts again": {
			inputSlice:  []int{5, 9, 7, 6, 1, 2},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"sorts more": {
			inputSlice:  []int{5, 7, 9, 6, 2, 1},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"works on single length slices": {
			inputSlice:  []int{5},
			outputSlice: []int{5},
		},
		"works on two length slices": {
			inputSlice:  []int{5, 1},
			outputSlice: []int{1, 5},
		},
	}

	for testName, test := range data {
		t.Run(testName, func(t *testing.T) {
			tc := test

			t.Parallel()

			algorithms.BubbleSort(tc.inputSlice)
			got := tc.inputSlice
			if !reflect.DeepEqual(got, tc.outputSlice) {
				t.Fatalf("BubbleSort did not sort correctly: got: %+v, want: %+v",
					got,
					tc.outputSlice,
				)
			}
			if !reflect.DeepEqual(got, tc.inputSlice) {
				t.Fatalf("BubbleSort did not mutate slice: %+v", tc.outputSlice)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	t.Parallel()

	data := map[string]struct {
		inputSlice  []int
		outputSlice []int
	}{
		"sorts": {
			inputSlice:  []int{9, 7, 5, 6, 1, 2},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"sorts again": {
			inputSlice:  []int{5, 9, 7, 6, 1, 2},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"sorts more": {
			inputSlice:  []int{5, 7, 9, 6, 2, 1},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"works on single length slices": {
			inputSlice:  []int{5},
			outputSlice: []int{5},
		},
		"works on two length slices": {
			inputSlice:  []int{5, 1},
			outputSlice: []int{1, 5},
		},
	}

	for testName, test := range data {
		t.Run(testName, func(t *testing.T) {
			tc := test

			t.Parallel()

			algorithms.SelectionSort(tc.inputSlice)
			got := tc.inputSlice
			if !reflect.DeepEqual(got, tc.outputSlice) {
				t.Fatalf("SelectionSort did not sort correctly: got: %+v, want: %+v",
					got,
					tc.outputSlice,
				)
			}
			if !reflect.DeepEqual(got, tc.inputSlice) {
				t.Fatalf("SelectionSort did not mutate slice: %+v", tc.outputSlice)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	t.Parallel()

	data := map[string]struct {
		inputSlice  []int
		outputSlice []int
	}{
		"sorts": {
			inputSlice:  []int{9, 5, 7, 6, 1, 2},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"sorts again": {
			inputSlice:  []int{5, 9, 7, 6, 1, 2},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"sorts more": {
			inputSlice:  []int{5, 7, 9, 6, 2, 1},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"works on single length slices": {
			inputSlice:  []int{5},
			outputSlice: []int{5},
		},
		"works on two length slices": {
			inputSlice:  []int{5, 1},
			outputSlice: []int{1, 5},
		},
	}

	for testName, test := range data {
		t.Run(testName, func(t *testing.T) {
			tc := test

			t.Parallel()

			algorithms.InsertionSort(tc.inputSlice)
			got := tc.inputSlice
			if !reflect.DeepEqual(got, tc.outputSlice) {
				t.Fatalf("InsertionSort did not sort correctly: got: %+v, want: %+v",
					got,
					tc.outputSlice,
				)
			}
			if !reflect.DeepEqual(got, tc.inputSlice) {
				t.Fatalf("InsertionSort did not mutate slice: %+v", tc.outputSlice)
			}
		})
	}
}

func TestQuicksort(t *testing.T) {
	t.Parallel()

	data := map[string]struct {
		inputSlice  []int
		outputSlice []int
	}{
		"sorts": {
			inputSlice:  []int{9, 5, 7, 6, 1, 2},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"sorts again": {
			inputSlice:  []int{5, 9, 7, 6, 1, 2},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"sorts more": {
			inputSlice:  []int{5, 7, 9, 6, 2, 1},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"doesn't get out of bounds errors": {
			inputSlice:  []int{9, 7, 6, 5, 2, 1},
			outputSlice: []int{1, 2, 5, 6, 7, 9},
		},
		"works on single length slices": {
			inputSlice:  []int{5},
			outputSlice: []int{5},
		},
		"works on two length slices": {
			inputSlice:  []int{5, 1},
			outputSlice: []int{1, 5},
		},
	}

	for testName, test := range data {
		t.Run(testName, func(t *testing.T) {
			tc := test

			t.Parallel()

			algorithms.Quicksort(tc.inputSlice, 0, len(tc.inputSlice)-1)
			got := tc.inputSlice
			if !reflect.DeepEqual(got, tc.outputSlice) {
				t.Fatalf("Quicksort did not sort correctly: got: %+v, want: %+v",
					got,
					tc.outputSlice,
				)
			}
			if !reflect.DeepEqual(got, tc.inputSlice) {
				t.Fatalf("Quicksort did not mutate slice: %+v", tc.outputSlice)
			}
		})
	}
}
