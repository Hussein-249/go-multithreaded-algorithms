/*
Incomplete implementation. Please ignore!
*/

package structures

import "math"

type RootishArray struct {
	blocks [][]interface{}
	size   int
}

func NewRootishArray() *RootishArray {
	// initilaize outer slice:  [][]interface{}, then make([]interface{}, 1)) to create an inner slice holding one element
	rootisharr := &RootishArray{
		blocks: [][]interface{}{make([]interface{}, 1)},
		size:   0,
	}

	return rootisharr
}

func (rootisharr *RootishArray) BlockIndex(index int) int {
	// math.Sqrt() accepts float64
	blockindex := int(math.Ceil(((math.Sqrt(float64(9 + (8 * index)))) - 3) / 2))
	return int(blockindex)
}

func (rootisharr *RootishArray) AddItem(item interface{}) bool {
	return false
}

// func (rootisharr *RootishArray) Size() int {
// 	s :=
// 	return 0
// }

// j=i−b(b+1)/2
// i is stored at location j within block b
