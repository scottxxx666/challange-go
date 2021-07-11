package _03nextgreater2

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	stack := make([]int, 0)
	for i := 2 * (len(nums) - 1); i >= 0; i-- {
		key := i % len(nums)
		for len(stack) > 0 && nums[key] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[key] = -1
		} else {
			res[key] = stack[len(stack)-1]
		}
		stack = append(stack, nums[key])
	}
	return res
}
