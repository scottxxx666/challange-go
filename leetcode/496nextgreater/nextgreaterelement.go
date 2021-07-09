package _96nextgreater

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	indexes := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		indexes[nums1[i]] = i
	}

	res := make([]int, len(nums1))
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if index, ok := indexes[nums2[i]]; ok {
			if len(stack) == 0 {
				res[index] = -1
			} else {
				res[index] = stack[len(stack)-1]
			}
		}
		stack = append(stack, nums2[i])
	}

	return res
}
