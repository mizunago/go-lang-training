package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (counter map[string]int) {
	words := strings.Fields(s)
	counter = make(map[string]int)
	for _, word := range words {
		count, check := counter[word]
		if check {
			counter[word] = count + 1
		} else {
			counter[word] = 1
		}
	}
	return
}

func main() {
	wc.Test(WordCount)
}
