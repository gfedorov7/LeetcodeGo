package tasks_from_leetcode

func TwoSum(nums []int, target int) []int {
	difs := make(map[int]int)

	for i, v := range nums {
		index := target - v
		if val, ok := difs[v]; ok {
			return []int{val, i}
		}
		difs[index] = i
	}

	return []int{}
}

////
