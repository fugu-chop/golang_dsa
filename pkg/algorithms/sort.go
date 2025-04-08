package algorithms

/*
BubbleSort takes an unsorted slice of int and returns that slice
with int sorted ascending (i.e. the original slice is mutated)
using a bubble sort algorithm.
*/
func BubbleSort(arr []int) []int {
	continueIterating := true
	end := len(arr) - 2

	for continueIterating {
		continueIterating = false
		for firstMarker := 0; firstMarker <= end; firstMarker++ {
			secondMarker := firstMarker + 1
			if arr[firstMarker] > arr[secondMarker] {
				arr[firstMarker], arr[secondMarker] = arr[secondMarker], arr[firstMarker]
				continueIterating = true
			}
		}
		end--
	}

	return arr
}

/*
SelectionSort takes an unsorted slice of int and returns that slice
with int sorted ascending (i.e. the original slice is mutated) using
the Selection sort algorithm.
*/
func SelectionSort(arr []int) []int {
	startIdx := 0

	for startIdx < len(arr) {
		lowestNumIdx := startIdx
		for i := startIdx; i < len(arr); i++ {
			if arr[i] < arr[lowestNumIdx] {
				lowestNumIdx = i
			}
		}
		if lowestNumIdx != startIdx {
			arr[startIdx], arr[lowestNumIdx] = arr[lowestNumIdx], arr[startIdx]
		}
		startIdx++
	}

	return arr
}

/*
InsertionSort takes an unsorted slice of int and returns that slice
with int sorted ascending (i.e. the original slice is mutated) using
the Insertion sort algorithm.
*/
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		insertionIdx := i - 1

		for insertionIdx >= 0 && arr[insertionIdx] > key {
			arr[insertionIdx+1] = arr[insertionIdx]
			insertionIdx -= 1
		}

		arr[insertionIdx+1] = key
	}

	return arr
}
