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

/*
Mergesort takes an unsorted slice of int and outputs a new slice
such that the ints are ascending using a merge sort algorithm
(i.e. the algorithm is non-mutating).
*/
func Mergesort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	first := Mergesort(items[len(items)/2:])
	second := Mergesort(items[:len(items)/2])

	return merge(first, second)
}

/*
merge takes two int slices and returns a new slice with the elements in ascending order.
It assumes that both `first` and `second` slices passed to it are in ascending order.
*/
func merge(first []int, second []int) []int {
	final := []int{}
	i := 0
	j := 0

	// Only iterate as far as the minimum length of the two slices
	// We have to keep two index variables as we may need to repeatedly
	// compare the same element.
	for i < len(first) && j < len(second) {
		if first[i] <= second[j] {
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}

	/*
		These are the 'leftover' elements if one slice is longer than the other
		Since both slices are already in ascending order, it is safe to append
		the remainder since they must be larger than all other elements.
	*/

	for ; i < len(first); i++ {
		final = append(final, first[i])
	}

	for ; j < len(second); j++ {
		final = append(final, second[j])
	}

	return final
}
