package _52meetingRoom

import "testing"

func Test_canJoinAllMeetings(t *testing.T) {
	type args struct {
		times [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{[][]int{}}, true},
		{"1 meeting", args{[][]int{{0, 9999}}}, true},
		{"2 overlap meeting", args{[][]int{{0, 1}, {1, 3}}}, false},
		{"2 overlap meeting", args{[][]int{{0, 2}, {1, 3}}}, false},
		{"2 overlap meeting", args{[][]int{{0, 3}, {1, 3}}}, false},
		{"2 overlap meeting", args{[][]int{{0, 4}, {1, 3}}}, false},
		{"2 meeting", args{[][]int{{0, 1}, {2, 3}}}, true},
		{"2 meeting", args{[][]int{{4, 5}, {2, 3}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJoinAllMeetings(tt.args.times); got != tt.want {
				t.Errorf("canJoinAllMeetings() = %v, want %v", got, tt.want)
			}
		})
	}
}
