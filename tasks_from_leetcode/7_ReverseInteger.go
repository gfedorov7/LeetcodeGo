package tasks_from_leetcode

import (
	"math"
	"strconv"
)

const DELIMITER int = 10

func reverse(x int) int {
	result := ""
	isSmaller := x < 0
	x = int(math.Abs(float64(x)))
	for x != 0 {
		part := x % DELIMITER
		x /= DELIMITER

		result = result + strconv.Itoa(part)
	}

	num, _ := strconv.Atoi(result)
	if float64(num) < -math.Pow(2, 31.0) || float64(num) > math.Pow(2, 31.0) {
		return 0
	}
	if isSmaller {
		num = num * -1
	}
	return num
}
