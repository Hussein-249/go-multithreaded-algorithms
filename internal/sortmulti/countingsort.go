package sortmulti

import "sync"

func CountingSort(inputArr []int, maximum int) []int {
	length := len(inputArr)

	if length == 0 {
		return []int{}
	}

	if length == 1 {
		return []int{inputArr[0]}
	}

	mid := length / 2
	start := 0
	end := length - 1
	freq := make([]int, maximum+1)
	res := make([]int, length)

	// use two goroutines to traverse list from both ends
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		for i := start; i < mid; i++ {
			freq[inputArr[i]]++
		}
	}()

	go func() {
		defer waitGroup.Done()
		for j := end; j >= mid; j-- {
			freq[inputArr[j]]++
		}
	}()

	waitGroup.Wait()

	// update freq to hold cumulative sum
	for i := 1; i < len(freq); i++ {
		freq[i] += freq[i-1]
	}

	// reverse traverse
	for i := end; i >= 0; i-- {
		res[freq[inputArr[i]]-1] = inputArr[i]
		freq[inputArr[i]]--
	}

	return res
}
