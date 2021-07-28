package _15kthlargestinarray

func findKthLargest(nums []int, k int) int {
	index := len(nums) - k
	i := 0
	j := len(nums) - 1
	for i <= j {
		temp := getIndex(nums, i, j)
		if temp == index {
			return nums[temp]
		} else if temp > index {
			j = temp - 1
		} else {
			i = temp + 1
		}
	}
	return -1
}

func getIndex(nums []int, start int, end int) int {
	pivot := nums[start]
	i := start + 1
	j := end
	for i <= j {
		if nums[i] >= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	nums[start], nums[j] = nums[j], nums[start]
	return i - 1
}
