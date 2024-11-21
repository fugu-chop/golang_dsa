package algorithms

/*
Bubble Sort takes an unsorted slice of int and returns
a copy of that slice with int sorted ascending (i.e. the original slice is unmodified).
*/
func BubbleSort(arr []int) []int {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	if len(arrCopy) == 1 {
		return arrCopy
	}

	swapped := true
	changes := 0

	for swapped {
		for firstMarker := 0; firstMarker <= len(arrCopy)-2; firstMarker++ {
			secondMarker := firstMarker + 1
			if arrCopy[firstMarker] > arrCopy[secondMarker] {
				temp := arrCopy[firstMarker]
				arrCopy[firstMarker] = arrCopy[secondMarker]
				arrCopy[secondMarker] = temp
				changes++
			}
		}
		if changes == 0 {
			swapped = false
		}
		changes = 0
	}

	return arrCopy
}
