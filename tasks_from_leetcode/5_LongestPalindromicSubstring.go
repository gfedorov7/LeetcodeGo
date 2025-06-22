package tasks_from_leetcode

func reverseString5(s string) string {
	newString := []rune(s)
	for i, j := 0, len(newString)-1; i < j; i, j = i+1, j-1 {
		newString[i], newString[j] = newString[j], newString[i]
	}
	return string(newString)
}

func getPalindrome(s, substr string, left, right int) string {
	if left > 0 && s[left-1:right+1] == reverseString5(s[left-1:right+1]) {
		return getPalindrome(s, s[left-1:right+1], left-1, right)
	}
	if left > 0 && right < len(s)-1 && s[left-1:right+2] == reverseString5(s[left-1:right+2]) {
		return getPalindrome(s, s[left-1:right+2], left-1, right+1)
	}
	return substr
}

func longestPalindrome(s string) string {
	currentLongestPalidrome := ""
	currentIndex := 0
	maxLongestPalidrome := ""

	if len(s) == 1 {
		return s
	}

	for currentIndex < len(s) {
		currentLongestPalidrome = getPalindrome(s, string(s[currentIndex]), currentIndex, currentIndex)
		if len(currentLongestPalidrome) > len(maxLongestPalidrome) {
			maxLongestPalidrome = currentLongestPalidrome
		}
		currentIndex++
	}

	return maxLongestPalidrome
}
