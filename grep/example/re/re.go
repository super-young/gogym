package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, _ := regexp.Compile("^a+b*$")

	fmt.Println(
		re.MatchString("ab"),
		re.MatchString("a"),
		re.MatchString("abbb"),
		re.MatchString("aaaaaabbb"),
		re.MatchString("ba"),
	)

	// true true true true false
}
