package timer

import (
	"fmt"
	"time"

	"go-multithreaded-algorithms/internal/sortmulti"
)

func TimeFunction() func() {

	startTime := time.Now().UnixNano()

	return func() {

		elapsedNanoseconds := time.Now().UnixNano() - startTime

		elapsedSeconds := float64(elapsedNanoseconds) / 1e9

		fmt.Printf("Time taken: %.9f nanoseconds\n", elapsedSeconds)

	}

}

func TimerWrapper() {

	var smallArray = []int{90, 56, 30, 23, 1, 65, 78, 89, 99, 45, 90, 56, 30, 23, 1, 65, 78, 89, 99, 45, 90, 56, 30, 23, 1, 65, 78, 89, 99, 45, 90, 56, 30, 23, 1, 65, 78, 89, 99, 45, 90, 56, 30, 23, 1, 65, 78, 89, 99, 45, 90, 56, 30, 23, 1, 65, 78, 89, 99, 45, 90, 56, 30, 23, 1, 65, 78, 89, 99, 45, 90, 56, 30, 23, 1, 65, 78, 89, 99, 45, 90, 56, 30, 23, 1, 65, 78, 89, 99, 45, 90, 56, 30, 23, 1, 65, 78, 89, 99, 45}

	defer TimeFunction()()

	sortmulti.MergeSort(smallArray)

}
