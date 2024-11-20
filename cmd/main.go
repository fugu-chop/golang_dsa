package main

import (
	"fmt"

	"github.com/fugu-chop/golang_dsa/pkg/algorithms"
)

func main() {
	sampleSlice := []int{0, 1, 2, 3, 5, 7, 9}

	idx := algorithms.BinarySearch(sampleSlice, 7)
	fmt.Println(idx)
}
