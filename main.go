package main // execute package

import (
	"fmt"

	"github.com/Hussein-249/go-multithreaded-algorithms/internal/sortmulti"
)

func main() {
	var smallArray = []int{90, 56, 30, 23, 1, 65, 78, 89, 99, 45}
	var emptyArray = []int{}
	fmt.Println(sortmulti.MergeSort(smallArray))
	fmt.Println(sortmulti.MergeSort(emptyArray)) // should return empty array
}
