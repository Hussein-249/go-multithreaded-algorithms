package sortmulti

// having trouble with constraints, defining this interface as a template to accept multiple data types
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func MergeSort[T Ordered](inputArr []T) []T {
	len := len(inputArr)

	// will see if bitwise ops more efficient,
	// but might not work since we accept floats
	if len == 0 {
		return []T{}
	}

	if len == 1 {
		return []T{inputArr[0]}
	}

	midIndex := len / 2
	leftSlice := inputArr[:midIndex]
	rightSlice := inputArr[midIndex:]

	// fmt.Println(leftSlice)
	// fmt.Println(rightSlice)

	MergeSort(leftSlice)
	MergeSort(rightSlice)

	tempSlice := []T{88}

	return tempSlice
}
