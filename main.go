package main // execute package

import (
	"fmt"

	"sort"

	"go-multithreaded-algorithms/internal/common"

	"go-multithreaded-algorithms/internal/sortmulti"

	"go-multithreaded-algorithms/internal/timer"
)

func main() {

	var smallArrayOdd = []int{90, 56, 30, 23, 1, 65, 78, 89, 99, 45, 44}

	fmt.Println(sortmulti.MergeSort(smallArrayOdd))

	timer.TimerWrapper()

	fmt.Println("Now sorting by in-built function:")

	var newArray = []int{90, 56, 30, 23, 1, 65, 78, 89, 99, 45}

	largeArray := common.RandArrGen(999999)

	sort.Ints(largeArray)

	sort.Ints(newArray)

	fmt.Println(newArray)

}
