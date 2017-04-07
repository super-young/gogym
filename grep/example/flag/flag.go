package main

import (
	"flag"
	"fmt"
)

var invertMatch = flag.Bool("v", false, "Selected lines are those not matching any of the specified patterns.")
var caseInsensitive = flag.Bool("i", false, "Perform case insensitive matching.")
var showLineNumber = flag.Bool("n", false, "Each output line is preceded by its relative line number in the file, starting at line 1.")

func main() {
	flag.Parse()
	args := flag.Args()

	fmt.Println("inverMatch", *invertMatch)
	fmt.Println("caseInsensitive", *caseInsensitive)
	fmt.Println("showLineNumber", *showLineNumber)
	fmt.Println("args", args)
}
