package tasks_from_leetcode

import (
	"sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	length := len(nums)

	if length%2 != 0 {
		return float64(nums[length/2])
	}

	return float64(nums[length/2]+nums[length/2-1]) / 2
}
