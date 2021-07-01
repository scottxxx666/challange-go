package _7insertInterval

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	for _, interval := range intervals {
		if newInterval[1] < interval[0] {
			res = append(res, newInterval)
			newInterval = interval
		} else if newInterval[0] > interval[1] {
			res = append(res, interval)
		} else {
			if interval[0] < newInterval[0] {
				newInterval[0] = interval[0]
			}
			if interval[1] > newInterval[1] {
				newInterval[1] = interval[1]
			}
		}
	}
	res = append(res, newInterval)

	return res
}
