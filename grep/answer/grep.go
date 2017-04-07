package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var invertMatch = flag.Bool("v", false, "Selected lines are those not matching any of the specified patterns.")
var caseInsensitive = flag.Bool("i", false, "Perform case insensitive matching.")
var showLineNumber = flag.Bool("n", false, "Each output line is preceded by its relative line number in the file, starting at line 1.")

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		flag.PrintDefaults()
		return
	}
	restr := args[0]

	// http://stackoverflow.com/questions/15326421/how-do-i-do-a-case-insensitive-regular-expression-in-go
	if *caseInsensitive {
		restr = "(?i)" + restr
	}

	re, err := regexp.Compile(restr)
	if err != nil {
		log.Fatalln(err)
	}

	if len(args) > 1 {
		files := args[1:]

		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			grep(filename, f, re)

			f.Close()
		}
	} else {
		grep("",os.Stdin, re)
	}
}

// https://golang.org/pkg/bufio/#example_Scanner_lines
func grep(filename string, r io.Reader, re *regexp.Regexp) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	lineno := 0

	for scanner.Scan() {
		lineno++
		txt := scanner.Text()
		test := re.MatchString(txt)

		if *invertMatch {
			test = !test
		}

		if test {
			if *showLineNumber {
				if filename != "" {
					fmt.Printf("%s:%d:", filename, lineno)
				} else {
					fmt.Printf("%d:", lineno)
				}
			}
			fmt.Println(txt)
		}
	}

	return scanner.Err()

}
