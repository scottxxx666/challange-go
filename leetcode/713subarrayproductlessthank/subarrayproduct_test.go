package _13subarrayproductlessthank

import "testing"

func Test_numSubarrayProductLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 num equal", args{[]int{1}, 1}, 0},
		{"1 num greater", args{[]int{1}, 2}, 1},
		{"2 num 2 solution", args{[]int{1, 2}, 3}, 3},
		{"2 num 2 solution", args{[]int{2, 3}, 4}, 2},
		{"2 num 1 solution", args{[]int{1, 3}, 3}, 1},
		{"3 num 6 solution", args{[]int{1, 2, 3}, 7}, 6},
		{"3 num 4 solution", args{[]int{1, 2, 3}, 6}, 4},
		{"4 num 7 solution", args{[]int{1, 2, 3, 4}, 7}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}
