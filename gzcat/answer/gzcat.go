package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		log.Println(err)
		return
	}
	r.Close()

	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		log.Println(err)
		return
	}
}
