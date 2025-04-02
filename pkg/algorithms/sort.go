package algorithms

/*
BubbleSort takes an unsorted slice of int and returns that slice
with int sorted ascending (i.e. the original slice is mutated) using a bubble sort algorithm.
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
with int sorted ascending (i.e. the original slice is mutated) using the Selection sort algorithm.
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
with int sorted ascending (i.e. the original slice is mutated) using the Insertion sort algorithm.
*/
func InsertionSort(arr []int) []int {
	var (
		originalStartIdx = 1
		startIdx         = originalStartIdx
	)

	for startIdx < len(arr) {
		for i := startIdx - 1; i >= 0; i-- {
			if arr[startIdx] < arr[i] {
				arr[startIdx], arr[i] = arr[i], arr[startIdx]
				startIdx = i
			} else {
				break
			}
		}

		originalStartIdx++
		startIdx = originalStartIdx
	}

	return arr
}
