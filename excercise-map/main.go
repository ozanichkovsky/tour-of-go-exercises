package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := map[string]int{}
	slice := strings.Fields(s)

	for _, word := range slice {
		result[word] += 1
	}

	return result
}

func main() {
	wc.Test(WordCount)
}

