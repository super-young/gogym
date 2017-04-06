package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for _, filename := range os.Args[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = f.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
