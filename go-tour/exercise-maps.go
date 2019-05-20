/*
	Go Tour Exercise: Maps https://tour.golang.org/moretypes/23
*/
package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	arrStrings := strings.Fields(s)
	fmt.Println(arrStrings)
	wordCountMap := make(map[string]int);

	for _, s := range arrStrings {
		wordCountMap[s] += 1;
	}
	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}