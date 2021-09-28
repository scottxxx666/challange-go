package _22sortarraybyparity2

func sortArrayByParityII(nums []int) []int {
	i, j := 1, 0
	for i < len(nums) && j < len(nums) {
		for i < len(nums) && nums[i]%2 == 1 {
			i += 2
		}
		for j < len(nums) && nums[j]%2 == 0 {
			j += 2
		}

		if i < len(nums) && j < len(nums) {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
