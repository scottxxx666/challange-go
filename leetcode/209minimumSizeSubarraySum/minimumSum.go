package _09minimumSizeSubarraySum

func minSubArrayLen(target int, nums []int) int {
	res := len(nums) + 1
	sum := 0
	for i, j := 0, 0; i <= j && j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			if j-i+1 < res {
				res = j - i + 1
			}
			sum -= nums[i]
			i++
		}
	}

	if res == len(nums)+1 {
		return 0
	}
	return res
}
