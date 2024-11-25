package main

import (
	"fmt"

	"github.com/fugu-chop/golang_dsa/pkg/algorithms"
)

func main() {
	sampleSlice := []int{0, 1, 2, 3, 5, 7, 9}

	idx := algorithms.BinarySearch(sampleSlice, 7)
	fmt.Println(idx)

	unsortedBubbleSortSlice := []int{2, 1, 3, 9, 7, 5}
	sortedBubbleSortSlice := algorithms.BubbleSort(unsortedBubbleSortSlice)
	fmt.Println(sortedBubbleSortSlice)

	unsortedSelectionSortSlice := []int{2, 1, 9, 7, 5, 3}
	sortedSelectionSortSlice := algorithms.SelectionSort(unsortedSelectionSortSlice)
	fmt.Println(sortedSelectionSortSlice)
}
