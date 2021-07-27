package _1firstMissingPositive

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Missing Positive", args{[]int{1, 3}}, 2},
		{"Missing One", args{[]int{2, 3}}, 1},
		{"Ignore Negative", args{[]int{1, 3, -1}}, 2},
		{"Ignore Zero", args{[]int{2, 3, 0}}, 1},
		{"No Missing", args{[]int{2, 3, 1}}, 4},
		{"Missing Swap", args{[]int{-1, 3, 1, 0}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
