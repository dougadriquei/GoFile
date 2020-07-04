package util

import (
	"log"
	"os"
)

//OpenFile read files on path
func OpenFile(path string) *os.File {
	csvfile, ferr := os.Open(path)
	if ferr != nil {
		log.Fatal(ferr)
	}
	return csvfile
}
