package utils

import (
	"fmt"
	"os"
)

// Check takes a potential error and panics if the error is defined.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Abort current application-execution.
func Abort() {
	fmt.Println("Aborting...")
	os.Exit(1)
}
