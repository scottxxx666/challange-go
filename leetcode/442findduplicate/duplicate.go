package _42findduplicate

import "math"

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		t := int(math.Abs(float64(nums[i]))) - 1
		if nums[t] > 0 {
			nums[t] *= -1
		} else if nums[t] < 0 {
			res = append(res, t+1)
		}
	}
	return res
}
