package main

import (
	"fmt"
)
// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
func main(){
	word1 := []string{"abc", "d", "defg"}
	word2 := []string{"abcddefg"}
	fmt.Print(arrayStringsAreEqual(word1, word2))
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    i := 0
	j := 0
	ii := 0
	jj := 0
	for i < len(word1) && j < len(word2) {
		if word1[i][ii] != word2[j][jj] {
			return false
		}
		ii += 1
		jj += 1
		if ii == len(word1[i]) {
			ii = 0
			i += 1
		}
		if jj == len(word2[j]) {
			jj = 0
			j += 1
		}
	}
	return i == len(word1) && j == len(word2)
}