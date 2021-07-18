package _79perfectsquare

import "math"

func numSquares(n int) int {
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s := math.Sqrt(float64(i))
		if s == float64(int(s)) {
			sums[i] = 1
			continue
		}
		temp := i
		for j := 1; j <= int(math.Ceil(float64(i/2))); j++ {
			if sums[j]+sums[i-j] < temp {
				temp = sums[j] + sums[i-j]
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
