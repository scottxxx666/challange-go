package _79perfectsquare

import "math"

func numSquares(n int) int {
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		temp := i
		for j := 1; j*j <= i; j++ {
			if i == j*j {
				temp = 1
			} else if 1+sums[i-j*j] < temp {
				temp = 1 + sums[i-j*j]
			}
		}
		sums[i] = temp
	}
	return sums[n]
}

func numSquares_bfs(n int) int {
	nums := make([]int, 0)
	for i := 1; i <= int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		nums = append(nums, i*i)
	}

	queue := []int{0}
	step := 0
	for {
		step++
		l := len(queue)
		for i := 0; i < len(nums); i++ {
			for j := 0; j < l; j++ {
				if nums[i]+queue[j] == n {
					return step
				}
				queue = append(queue, nums[i]+queue[j])
			}
		}
	}
}
