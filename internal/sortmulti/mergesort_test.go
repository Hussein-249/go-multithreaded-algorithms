/*
Testfile targeting the mergesort algorithm, source mergesort.go

To run the tests in this file only, execute
go test -run .\mergesort_test.go

Run all tests in the package via
go test -count=1 ./...
*/

// ignore VS Code suggestions - do not include _test in package name
// cannot export these functions if _test.go file?
package sortmulti

// must import testify/assert although the modules listed in go.mod is stretchr/testify
import (
	"os"

	"sort"

	"testing"

	"algoroutine/internal/common"

	"github.com/stretchr/testify/assert"
)

func LoadTextFileData(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// flie.CLose().Error() will handle the close errorseparately from code
	// file.Close() returns error if it occurs, up to programmer to handle it
	defer file.Close()
}

func TestMergeSort(t *testing.T) {

	sampleArr := common.RandArrGenFloat64(1000000, -600.0, 99999999999.9)
	expectedArr := make([]float64, len(sampleArr))
	copy(expectedArr, sampleArr)
	sort.Float64s(expectedArr)

	result := MergeSort(sampleArr)

	assert.Equal(t, expectedArr, result, "Should correctly sort.")
}

func TestMergeSortInts(t *testing.T) {

	sampleArr := common.RandArrGen(1000000, 700000)

	expectedArr := make([]int, len(sampleArr))

	copy(expectedArr, sampleArr)

	sort.Ints(expectedArr)

	result := MergeSort(sampleArr)

	assert.Equal(t, expectedArr, result, "Should correctly sort, MergeSort not sorting!")
}
