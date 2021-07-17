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
