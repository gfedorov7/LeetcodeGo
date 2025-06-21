package tasks_from_leetcode

import (
	"math"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	num, err := strconv.Atoi(s)
	if err != nil {
		lastNum := len(s)
		for idx, char := range s {
			if idx == 0 && (string(char) == "-" || string(char) == "+") {
				continue
			}
			if !strings.Contains("0123456789", string(char)) {
				lastNum = idx
				break
			}
		}
		num, _ = strconv.Atoi(s[:lastNum])
	}

	if v := -math.Pow(2.0, 31.0); float64(num) < v {
		return int(v)
	}
	if v := math.Pow(2.0, 31) - 1; float64(num) > v {
		return int(v)
	}

	return num
}
