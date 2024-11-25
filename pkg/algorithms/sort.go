package algorithms

/*
Bubble Sort takes an unsorted slice of int and returns that slice
with int sorted ascending (i.e. the original slice is mutated).
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
Selection Sort takes an unsorted slice of int and returns that slice
with int sorted ascending (i.e. the original slice is mutated).
*/
func SelectionSort(arr []int) []int {
	startIdx := 0
	arrLen := len(arr) - 1

	for startIdx < len(arr) {
		lowestNumIdx := startIdx
		for i := startIdx; i <= arrLen; i++ {
			if arr[i] < arr[lowestNumIdx] {
				lowestNumIdx = i
			}
		}
		arr[startIdx], arr[lowestNumIdx] = arr[lowestNumIdx], arr[startIdx]
		startIdx++
	}

	return arr
}
