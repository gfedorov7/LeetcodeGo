package tasks_from_leetcode

import "strconv"

func reverseString(s string) string {
	newString := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		newString[i], newString[j] = newString[j], newString[i]
	}
	return string(newString)
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	return reverseString(s) == s
}
