package tasks_from_leetcode

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	matrix := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]string, 0)
	}

	isReduce := false
	numOfRow := 0
	for i := 0; i < len(s); i++ {
		matrix[numOfRow] = append(matrix[numOfRow], string(s[i]))

		if isReduce {
			numOfRow--
			if numOfRow == 0 {
				isReduce = false
			}
		} else {
			numOfRow++
			if numOfRow == numRows-1 {
				isReduce = true
			}
		}
	}

	result := ""
	for _, row := range matrix {
		result += strings.Join(row, "")
	}
	return result
}
