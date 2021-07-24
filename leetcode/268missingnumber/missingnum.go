package _68missingnumber

func missingNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if i == nums[i] || i == len(nums) || nums[i] == len(nums) {
			i++
		} else {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}
