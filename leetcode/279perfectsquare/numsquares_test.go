package _79perfectsquare

import "testing"

func Test_numSquares(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 1},
		{"2", args{2}, 2},
		{"3", args{3}, 3},
		{"4", args{4}, 1},
		{"5", args{5}, 2},
		{"6", args{6}, 3},
		{"7", args{7}, 4},
		{"8", args{8}, 2},
		{"9", args{9}, 1},
		{"11", args{11}, 3},
		{"12", args{12}, 3},
		{"13", args{13}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.args.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
