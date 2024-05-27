/*
Testfile targeting the rootisharray data structure, source rootisharray.go

To run the tests in this file only, execute
go test -run .\rootisharray_test.go

Run all tests in the package via
go test -count=1 ./...
*/
package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockIndex(t *testing.T) {
	ra := NewRootishArray()
	result := ra.BlockIndex(15)
	assert.Equal(t, 5, result, "BlockIndex functions should return correct block index (15 -> 5) based on element position in entire structure.")
}
