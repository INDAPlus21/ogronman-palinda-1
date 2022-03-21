package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	chars := strings.Fields(s)
	wordcount := make(map[string]int)

	for i := range chars {
		wordcount[chars[i]]++
	}
	return wordcount
}

func main() {
	wc.Test(WordCount)
}
