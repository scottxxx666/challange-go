package _6threeSumClosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(sum-target)) < math.Abs(float64(res-target)) {
				res = sum
			}

			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
