package _00numberofislands

import "testing"

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1*1 1 island", args{[][]byte{{1}}}, 1},
		{"1*1 0 island", args{[][]byte{{0}}}, 0},
		{"2*2 1 island", args{[][]byte{{0, 1}, {0, 1}}}, 1},
		{"2*2 1 island", args{[][]byte{{1, 1}, {0, 0}}}, 1},
		{"2*2 2 island", args{[][]byte{{1, 0}, {0, 1}}}, 2},
		{"2*2 1 island", args{[][]byte{{1, 1}, {0, 1}}}, 1},
		{"2*2 1 island", args{[][]byte{{1, 0}, {1, 1}}}, 1},
		{"1*2 1 island", args{[][]byte{{1, 0}}}, 1},
		{"2*1 1 island", args{[][]byte{{1}, {1}}}, 1},
		{"5*5 1 island", args{[][]byte{
			{1, 0, 1, 0, 1},
			{1, 1, 1, 1, 0},
			{0, 0, 0, 1, 0},
			{1, 0, 1, 1, 0},
			{0, 1, 1, 0, 0},
		}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
