package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var invertMatch = flag.Bool("v", false, "Selected lines are those not matching any of the specified patterns.")
var caseInsensitive = flag.Bool("i", false, "Perform case insensitive matching.")
var showLineNumber = flag.Bool("n", false, "Each output line is preceded by its relative line number in the file, starting at line 1.")

func main() {
	flag.Parse()
	pattern := regexp.QuoteMeta(flag.Args()[0])

	var inputSource io.Reader
	if len(flag.Args()) > 1 {
		file, err := os.Open(flag.Args()[1])
		if err != nil {
			fmt.Println("failed to open file")
			return
		}
		defer file.Close()

		inputSource = file
	} else {
		inputSource = os.Stdin
	}

	if *caseInsensitive {
		pattern = "(?i)" + pattern
	}

	re, _ := regexp.Compile(pattern)

	scanner := bufio.NewScanner(inputSource)

	for i := 0; scanner.Scan(); i++ {
		lineNo := ""
		text := scanner.Text()
		if *showLineNumber {
			lineNo = strconv.Itoa(i) + ":"
		}
		if *invertMatch != re.MatchString(text) {
			fmt.Println(lineNo, text)
		}
	}
}
