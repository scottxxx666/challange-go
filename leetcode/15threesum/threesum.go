package _5threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) <= 2 {
		return res
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			if nums[j]+nums[k] == -nums[i] {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j-1] == nums[j] {
					j++
				}
				for k > j && nums[k+1] == nums[k] {
					k--
				}
			} else if nums[j]+nums[k] > -nums[i] {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
