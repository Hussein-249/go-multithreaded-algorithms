package main // execute package

import (
	"fmt"

	"sort"

	"algoroutine/internal/common"

	"algoroutine/internal/sortmulti"

	"algoroutine/internal/timer"
)

func main() {

	var smallArrayOdd = []int{90, 56, 30, 23, 1, 65, 78, 89, 99, 45, 44}

	fmt.Println(sortmulti.MergeSort(smallArrayOdd))

	timer.TimerWrapper(sortmulti.MergeSort[int], smallArrayOdd)

	fmt.Println("Now sorting by in-built function:")

	largeArray := common.RandArrGen(999999, 19999999999)

	sort.Ints(largeArray)

	fmt.Println(largeArray)

}
