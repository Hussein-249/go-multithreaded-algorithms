package sortmulti

import (
	"sort"

	"testing"

	"go-multithreaded-algorithms/internal/common"

	"github.com/stretchr/testify/assert"
)

func TestCountingSort(t *testing.T) {

	sampleArr := common.RandArrGen(10000, 100)

	expectedArr := make([]int, len(sampleArr))

	copy(expectedArr, sampleArr)

	sort.Ints(expectedArr)

	maximum := 0

	for _, value := range sampleArr {
		if value > maximum {
			maximum = value
		}
	}

	result := CountingSort(sampleArr, maximum)

	assert.Equal(t, expectedArr, result, "Should correctly sort.")
}
