package common

import (
	"math/rand"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// for later, passing as pointer / addr
// type MergeFunc func[T Ordered]([]T)
// []T

func RandArrGen(size int) []int {

	max := 1000000

	results := make([]int, size)

	for i := 0; i < size; i++ {

		randomIntInRange := rand.Intn(max)

		results[i] = randomIntInRange
	}
	return results
}
