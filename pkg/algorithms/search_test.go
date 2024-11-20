package algorithms_test

import (
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/algorithms"
)

func TestBinarySearch(t *testing.T) {
	data := map[string]struct {
		inputSlice []int
		inputInt   int
		wantIdx    int
	}{
		"number exists": {
			inputSlice: []int{1, 3, 5, 7, 9},
			inputInt:   3,
			wantIdx:    1,
		},
		"number doesn't exist": {
			inputSlice: []int{1, 3, 5, 7, 9},
			inputInt:   10,
			wantIdx:    -1,
		},
	}

	for testCase, testData := range data {
		t.Run(testCase, func(t *testing.T) {
			got := algorithms.BinarySearch(testData.inputSlice, testData.inputInt)
			if got != testData.wantIdx {
				t.Fatalf("BinarySearch index does not match: got: %d, want: %d",
					got,
					testData.wantIdx,
				)
			}
		})
	}
}
