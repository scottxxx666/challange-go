package _86wallsandgates

func nearestGate(m [][]int) {
	var queue [][]int
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if m[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	step := 0
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			x, y := queue[0][0], queue[0][1]
			if m[x][y] == 2147483647 {
				m[x][y] = step
			}
			if x+1 < len(m) && m[x+1][y] == 2147483647 {
				queue = append(queue, []int{x + 1, y})
			}
			if x > 1 && m[x-1][y] == 2147483647 {
				queue = append(queue, []int{x - 1, y})
			}
			if y+1 < len(m[0]) && m[x][y+1] == 2147483647 {
				queue = append(queue, []int{x, y + 1})
			}
			if y > 1 && m[x][y-1] == 2147483647 {
				queue = append(queue, []int{x, y - 1})
			}
		}
		queue = queue[l:]
		step++
	}
}
