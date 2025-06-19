package tasks_from_leetcode

import (
	"strings"
)

func LengthOfLongestSubstring(s string) int {
	l, r := 0, 1
	mx := 1

	if len(s) < 1 {
		return 0
	}
	for {
		if mx >= r-l && r+1 > len(s) {
			break
		}

		if strings.Count(s[l:r], string(s[r])) > 0 {
			l++
			continue
		} else {
			r++
		}

		if r-l > mx {
			mx = r - l
		}
	}
	return mx
}
