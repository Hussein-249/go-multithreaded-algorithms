/*
Testfile targeting the countingsort algorithm, source countingsort.go

To run the tests in this file only, execute
go test -run .\countingsort_test.go

Run all tests in the package via
go test -count=1 ./...
*/

// must import testify/assert although the modules listed in go.mod is stretchr/testify
package sortmulti

import (
	"sort"

	"testing"

	"algoroutine/internal/common"

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
