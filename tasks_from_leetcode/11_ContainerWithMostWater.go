package tasks_from_leetcode

func minFromTwoNumber(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	maxAreaInContainer := 0
	left, right := 0, len(height)-1

	for left < right {
		currentArea := minFromTwoNumber(height[left], height[right]) * (right - left)
		if currentArea > maxAreaInContainer {
			maxAreaInContainer = currentArea
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxAreaInContainer
}
