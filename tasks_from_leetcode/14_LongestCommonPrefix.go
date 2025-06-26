package tasks_from_leetcode

import (
	"fmt"
)

func findSmallestWord(strs []string) string {
	smallestWord := strs[0]
	smallestWordLen := len(smallestWord)
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < smallestWordLen {
			smallestWord = strs[i]
			smallestWordLen = len(smallestWord)
		}
	}
	return smallestWord
}

func longestCommonPrefix(strs []string) string {
	smallestWord := findSmallestWord(strs)
	for j := len(smallestWord); j > 0; j-- {
		contains := true
		for _, word := range strs {
			fmt.Println(word, smallestWord[:j])
			if prefix := smallestWord[:j]; !(word[:j] == prefix) {
				contains = false
			}
		}

		if contains {
			return smallestWord[:j]
		}
	}
	return ""
}
