package _77squareofasortedarray

import "math"

func sortedSquares(nums []int) []int {
	res := make([]int, 0)
	for i, j := 0, len(nums)-1; i <= j; {
		if math.Pow(float64(nums[i]), 2) >= math.Pow(float64(nums[j]), 2) {
			res = append(res, int(math.Pow(float64(nums[i]), 2)))
			i++
		} else {
			res = append(res, int(math.Pow(float64(nums[j]), 2)))
			j--
		}
	}
	for i, j := 0, len(nums)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
