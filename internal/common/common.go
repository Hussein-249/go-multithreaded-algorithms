package common

import (
	"math/rand"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// for later, passing as pointer / addr
type SortSignature[T Ordered] func([]T) []T

func RandArrGen(size, max int) []int {
	// Returns a slice of randomly generated integers

	results := make([]int, size)

	for i := 0; i < size; i++ {

		randomIntInRange := rand.Intn(max)

		results[i] = randomIntInRange
	}
	return results
}

func RandArrGenFloat64(size int, min, max float64) []float64 {

	results := make([]float64, size)

	for i := 0; i < size; i++ {
		randf := rand.Float64()
		randfs := randf*(max-min) + min
		results[i] = randfs
	}

	return results
}
