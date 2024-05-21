package testdata

import (
	"fmt"
	"math/rand"
)

func randomNumberGenerator() {
	min := 10
	max := 20
	randomIntInRange := rand.Intn(max-min+1) + min
	fmt.Println("Random Integer in Range:", randomIntInRange)
}
