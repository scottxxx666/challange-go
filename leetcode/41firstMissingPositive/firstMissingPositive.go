package _1firstMissingPositive

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num > 0 && num < len(nums) && nums[num-1] != num {
			a := nums[num-1]
			nums[num-1] = num
			nums[i] = a
			i--
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
