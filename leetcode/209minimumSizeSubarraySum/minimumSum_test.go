package _09minimumSizeSubarraySum

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 element 0 solution", args{2, []int{1}}, 0},
		{"1 element 1 solution", args{1, []int{1}}, 1},
		{"2 element result 2", args{3, []int{1, 2}}, 2},
		{"2 element result 1", args{2, []int{1, 2}}, 1},
		{"3 element result 1", args{3, []int{1, 2, 3}}, 1},
		{"3 element result 2", args{5, []int{1, 2, 3}}, 2},
		{"4 element result 2", args{5, []int{4, 0, 2, 3}}, 2},
		{"4 element result 2", args{7, []int{4, 1, 2, 5}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
