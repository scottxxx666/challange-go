package _52openthelock

import "strconv"

func openLock(deadEnds []string, target string) int {
	queue := []string{"0000"}
	steps := 0

	m := make(map[string]bool)
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			temp := queue[0]
			queue = queue[1:]

			if temp == target {
				return steps
			}

			m[temp] = true
			for j := 0; j < 4; j++ {
				for k := -1; k <= 1; k += 2 {
					atoi, _ := strconv.Atoi(string(temp[j]))
					t := temp[0:j] + strconv.Itoa((atoi+k+10)%10) + temp[j+1:4]
					if !m[t] && !isDead(deadEnds, t) {
						queue = append(queue, t)
					}
				}
			}
		}
		steps++
	}
	return -1
}

func isDead(deadEnd []string, target string) bool {
	for _, s := range deadEnd {
		if s == target {
			return true
		}
	}
	return false
}
