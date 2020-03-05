package util

import (
	"fmt"
	"log"
	"os"
)

// TryOpen tries to open a given file in multiple location
// ie. running main.go from cmd or from root.
func TryOpen(path string) *os.File {
	var f *os.File
	f, err := os.Open(fmt.Sprintf("../%s", path))
	if err != nil {
		f, err = os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		return f

	}
	return f

}
