package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	pow := strings.Fields(s)
	mapa := make(map[string]int)
	for i := range pow {
		mapa[pow[i]] = 1
	}
	return mapa
}

func main() {
	wc.Test(WordCount)
}
