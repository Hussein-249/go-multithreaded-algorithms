// ignore VS Code suggestions - do not include _test in package name
// cannot export these functions if _test.go file
package sortmulti

import (
	"os"
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
