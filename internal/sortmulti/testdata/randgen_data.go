package testdata

import (
	"fmt"
	"math/rand"
	"os"
)

func randomNumberGenerator() {
	min := 10
	max := 20
	randomIntInRange := rand.Intn(max-min+1) + min
	fmt.Println("Random Integer in Range:", randomIntInRange)
}

func readTextFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// flie.CLose().Error() will handle the close errorseparately from code
	// file.Close() returns error if it occurs, up to programmer to handle it
	defer file.Close()
}
