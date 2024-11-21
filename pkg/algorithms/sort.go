package algorithms

/*
Bubble Sort takes an unsorted slice of int and returns
a copy of that slice with int sorted ascending (i.e. the original slice is unmodified).
*/
func BubbleSort(arr []int) []int {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)

	continueIterating := true
  end := len(arrCopy)-2

	for continueIterating {
    continueIterating = false
		for firstMarker := 0; firstMarker <= end; firstMarker++ {
			secondMarker := firstMarker + 1
			if arrCopy[firstMarker] > arrCopy[secondMarker] {
				arrCopy[firstMarker], arrCopy[secondMarker] = arrCopy[secondMarker], arrCopy[firstMarker]
				continueIterating = true
			}
		}
    end--
	}

	return arrCopy
}
