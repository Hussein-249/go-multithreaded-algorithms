package sortmulti

import (
	"sync"
)

// having trouble with constraints, defining this interface as a template to accept multiple data types
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func MergeSort[T Ordered](inputArr []T) []T {
	// do not use len - shadows len function!
	length := len(inputArr)

	// will see if bitwise ops more efficient,
	// but might not work since we accept floats
	if length == 0 {
		return []T{}
	}

	if length == 1 {
		return []T{inputArr[0]}
	}

	midIndex := length / 2
	leftSlice := inputArr[:midIndex]
	rightSlice := inputArr[midIndex:]

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	var lower, upper []T

	go func() {
		defer waitGroup.Done()
		lower = MergeSort(leftSlice)
	}()

	go func() {
		defer waitGroup.Done()
		upper = MergeSort(rightSlice)
	}()

	waitGroup.Wait()

	return merge(lower, upper)
}

func merge[T Ordered](lower []T, upper []T) []T {
	var results = make([]T, len(lower)+len(upper))
	i, j, k := 0, 0, 0

	for i < len(lower) && j < len(upper) {
		if lower[i] < upper[j] {
			results[k] = lower[i]
			i++
		} else {
			results[k] = upper[j]
			j++
		}
		k++
	}

	for i < len(lower) {
		results[k] = lower[i]
		k++
		i++
	}

	for j < len(upper) {
		results[k] = upper[j]
		j++
		k++
	}

	return results
}
