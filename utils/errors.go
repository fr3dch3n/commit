package utils

import (
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abort() {
	fmt.Println("Aborting...")
	os.Exit(1)
}
