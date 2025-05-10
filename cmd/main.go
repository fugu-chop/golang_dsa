package main

import (
	"fmt"

	"github.com/fugu-chop/golang_dsa/pkg/algorithms"
)

func main() {
	sampleSlice := []int{0, 1, 2, 3, 5, 7, 9}
	idx := algorithms.BinarySearch(sampleSlice, 7)
	fmt.Println(idx)

	bubbleSortSlice := []int{2, 1, 3, 9, 7, 5}
	algorithms.BubbleSort(bubbleSortSlice)
	fmt.Println(bubbleSortSlice)

	selectionSortSlice := []int{2, 1, 9, 7, 5, 3}
	algorithms.SelectionSort(selectionSortSlice)
	fmt.Println(selectionSortSlice)

	insertionSortSlice := []int{2, 1, 9, 7, 5, 3}
	algorithms.InsertionSort(insertionSortSlice)
	fmt.Println(insertionSortSlice)

	unsortedQuicksortSlice := []int{2, 1, 9, 7, 5, 3}
	algorithms.Quicksort(unsortedQuicksortSlice, 0, len(unsortedQuicksortSlice)-1)
	fmt.Println(unsortedQuicksortSlice)
}
