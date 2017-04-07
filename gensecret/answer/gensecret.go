package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
	"strconv"
)

func main() {
	nbytes, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	buf := make([]byte, nbytes)

	n, err := rand.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	if n != nbytes {
		log.Fatalf("Could not read %d random bytes\n", nbytes)
	}

	output := os.Stdout

	enc := base64.NewEncoder(base64.StdEncoding, output)
	_, err = enc.Write(buf)
	if err != nil {
		log.Println(err)
	}
	enc.Close()

	output.WriteString("\n")
}
