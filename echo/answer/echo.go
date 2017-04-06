package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		fmt.Print(arg)
		if i != len(os.Args)-1 {
			fmt.Print(" ")
		}
	}

	fmt.Print("\n")
}
