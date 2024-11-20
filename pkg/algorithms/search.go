package algorithms

/*
BinarySearch takes a sorted slice of ints (ascending) and an int to
search for. It returns an int, representing the index of the slice
where the int to search for is found. If the int is not found, -1 is
returned.
*/
func BinarySearch(sortedSlice []int, searchElement int) int {
	startIdx := 0
	endIdx := len(sortedSlice) - 1

	for startIdx <= endIdx {
		midpointIdx := (endIdx + startIdx) / 2
		foundValue := sortedSlice[midpointIdx]
		switch {
		case foundValue == searchElement:
			return midpointIdx
		case foundValue > searchElement:
			endIdx = midpointIdx - 1
		case foundValue < searchElement:
			startIdx = midpointIdx + 1
		}
	}

	return -1
}
