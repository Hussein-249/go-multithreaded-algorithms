package sortmulti

func CountingSort[T Ordered](inputArr []T) []T {
	length := len(inputArr)

	if length == 0 {
		return []T{}
	}

	if length == 1 {
		return []T{inputArr[0]}
	}

	return []T{}
}
