package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// Exercise: Maps
// Implement WordCount. It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.
// You might find strings.Fields helpful.

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	count := make(map[string]int)
	for f := range fields {
		v, ok := count[fields[f]]
		if ok {
			count[fields[f]] = v + 1
		} else {
			count[fields[f]] = 1
		}
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
