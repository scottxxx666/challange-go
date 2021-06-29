package _52meetingRoom

import "sort"

func canJoinAllMeetings(times [][]int) bool {
	sort.SliceStable(times, func(i, j int) bool {
		return times[i][0] < times[j][0]
	})

	for i := 0; i < len(times)-1; i++ {
		if times[i][1] >= times[i+1][0] {
			return false
		}
	}
	return true
}
