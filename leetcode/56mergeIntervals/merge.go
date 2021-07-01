package _6mergeIntervals

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	len := len(intervals)
	if len <= 1 {
		return intervals
	}

	res := make([][]int, 0)
	for i := 1; i < len; i++ {
		if intervals[i][0] <= intervals[i-1][1] {
			intervals[i][0] = intervals[i-1][0]
			if intervals[i-1][1] > intervals[i][1] {
				intervals[i][1] = intervals[i-1][1]
			}
		} else {
			res = append(res, []int{intervals[i-1][0], intervals[i-1][1]})
		}
	}
	res = append(res, []int{intervals[len-1][0], intervals[len-1][1]})
	return res
}
