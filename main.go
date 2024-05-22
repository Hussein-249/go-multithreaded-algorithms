package main // execute package

import (
	"fmt"

	"sort"

	"go-multithreaded-algorithms/internal/sortmulti"

	"go-multithreaded-algorithms/internal/timer"
)

func main() {
	var smallArrayOdd = []int{90, 56, 30, 23, 1, 65, 78, 89, 99, 45, 44}

	fmt.Println(sortmulti.MergeSort(smallArrayOdd))

	timer.TimerWrapper()

	fmt.Println("Now sorting by in-built function:")

	var newArray = []int{90, 56, 30, 23, 1, 65, 78, 89, 99, 45}

	sort.Ints(newArray)

	fmt.Println(newArray)

}
