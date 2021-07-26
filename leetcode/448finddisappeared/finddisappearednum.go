package _48finddisappeared

func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		if nums[i] != (i+1) && (nums[nums[i]-1] != nums[i]) {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
