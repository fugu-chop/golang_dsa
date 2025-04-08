package algorithms_test

import (
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/algorithms"
)

func TestBinarySearch(t *testing.T) {
	t.Parallel()

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

	for testCase, test := range data {
		t.Run(testCase, func(t *testing.T) {
			tc := test

			t.Parallel()

			got := algorithms.BinarySearch(tc.inputSlice, tc.inputInt)
			if got != tc.wantIdx {
				t.Fatalf("BinarySearch index does not match: got: %d, want: %d",
					got,
					tc.wantIdx,
				)
			}
		})
	}
}
