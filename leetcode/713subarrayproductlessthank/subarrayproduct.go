package _13subarrayproductlessthank

func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	product := 1

	for i, j := 0, 0; i <= j && j < len(nums); {
		product *= nums[j]
		if product < k && j < len(nums) {
			res += j - i + 1
			j++
		} else {
			product = product / nums[i] / nums[j]
			i++
		}
	}
	return res
}
