package algorithms_test

import (
	"reflect"
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/algorithms"
)

func TestBubbleSort(t *testing.T) {
	data := map[string]struct {
		inputSlice  []int
		outputSlice []int
	}{
		"sorts": {
			inputSlice:  []int{9, 7, 5, 6, 1, 2},
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
			got := algorithms.BubbleSort(test.inputSlice)
			if !reflect.DeepEqual(got, test.outputSlice) {
				t.Fatalf("BubbleSort did not sort correctly: got: %+v, want: %+v",
					got,
					test.outputSlice,
				)
			}
			if !reflect.DeepEqual(got, test.inputSlice) {
				t.Fatalf("BubbleSort did not mutate slice: %+v", test.outputSlice)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	data := map[string]struct {
		inputSlice  []int
		outputSlice []int
	}{
		"sorts": {
			inputSlice:  []int{9, 7, 5, 6, 1, 2},
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
			got := algorithms.SelectionSort(test.inputSlice)
			if !reflect.DeepEqual(got, test.outputSlice) {
				t.Fatalf("SelectionSort did not sort correctly: got: %+v, want: %+v",
					got,
					test.outputSlice,
				)
			}
			if !reflect.DeepEqual(got, test.inputSlice) {
				t.Fatalf("SelectionSort did not mutate slice: %+v", test.outputSlice)
			}
		})
	}
}
