package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	jsonfile := os.Args[1]
	f, err := os.Open(jsonfile)
	if err != nil {
		log.Fatalln(err)
	}

	var data interface{}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("decoded data: %v", data)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	err = enc.Encode(data)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println("data: %v", data)
}
