package timer

import (
	"fmt"
	"time"

	"go-multithreaded-algorithms/internal/common"
)

func TimeFunction() func() {

	startTime := time.Now().UnixNano()

	return func() {

		elapsedNanoseconds := time.Now().UnixNano() - startTime

		elapsedSeconds := float64(elapsedNanoseconds) / 1e9

		fmt.Printf("Time taken: %.9f nanoseconds\n", elapsedSeconds)

	}

}

func TimerWrapper[T common.Ordered](targetFunc common.SortSignature[T], data []T) {

	defer TimeFunction()()

	targetFunc(data)

}
