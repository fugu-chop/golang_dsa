package algorithms

/*
BubbleSort takes an unsorted slice of int and mutates that
slice such that the ints are ascending using a bubble sort algorithm.
*/
func BubbleSort(arr []int) {
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
}

/*
SelectionSort takes an unsorted slice of int and mutates that
slice such that the ints are ascending using a selection sort algorithm.
*/
func SelectionSort(arr []int) {
	for startIdx := range arr {
		lowestNumIdx := startIdx
		for i := startIdx + 1; i < len(arr); i++ {
			if arr[i] < arr[lowestNumIdx] {
				lowestNumIdx = i
			}
		}
		if lowestNumIdx != startIdx {
			arr[startIdx], arr[lowestNumIdx] = arr[lowestNumIdx], arr[startIdx]
		}
	}
}

/*
InsertionSort takes an unsorted slice of int and mutates that
slice such that the ints are ascending using an insertion sort algorithm.
*/
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		insertionIdx := i - 1

		for insertionIdx >= 0 && arr[insertionIdx] > key {
			arr[insertionIdx+1] = arr[insertionIdx]
			insertionIdx -= 1
		}

		arr[insertionIdx+1] = key
	}
}

/*
Quicksort takes an unsorted slice of int and uses partitioning
to sort the slice with int ascending using the Quicksort algorithm.
*/
func Quicksort(arr []int, leftIdx int, rightIdx int) {
	if rightIdx-leftIdx <= 0 {
		return
	}

	pivotIdx := partition(arr, leftIdx, rightIdx)
	Quicksort(arr, leftIdx, pivotIdx-1)
	Quicksort(arr, pivotIdx+1, rightIdx)
}

/*
`partition` mutates a slice by setting a pivot index (one position to the
left of rightIdx) and ensuring all elements to the left of the pivot
index are lower than the pivot value by swapping values (though not
necessarily sorted).

This ensures that the pivot is in the correct position after an iteration
of the `partition` function.

`partition` is intended to be called recursively as part of a Quicksort algorithm.
*/
func partition(arr []int, leftIdx int, rightIdx int) int {
	pivotIdx := rightIdx
	pivot := arr[pivotIdx]
	rightIdx -= 1

	for {
		/*
			This doesn't need an OOB guard as arr[leftIdx]
			will never exceed pivot in an already sorted
			ascending slice
		*/
		for arr[leftIdx] < pivot {
			leftIdx++
		}

		/*
			Requires an OOB guard as in a reverse sorted
			slice, arr[rightIdx] will always be greater
			than pivot
		*/
		for arr[rightIdx] > pivot && rightIdx > 0 {
			rightIdx--
		}

		if leftIdx >= rightIdx {
			break
		}

		arr[leftIdx], arr[rightIdx] = arr[rightIdx], arr[leftIdx]
		leftIdx++
	}

	arr[leftIdx], arr[pivotIdx] = arr[pivotIdx], arr[leftIdx]

	return leftIdx
}
