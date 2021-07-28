package _15kthlargestinarray

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 num", args{[]int{2}, 1}, 2},
		{"2 num", args{[]int{2, 3}, 1}, 3},
		{"2 num", args{[]int{2, 3}, 2}, 2},
		{"2 num", args{[]int{6, 5}, 1}, 6},
		{"2 num same", args{[]int{6, 6}, 1}, 6},
		{"2 num same", args{[]int{6, 6}, 2}, 6},
		{"3 num", args{[]int{5, 6, 7}, 1}, 7},
		{"3 num", args{[]int{5, 6, 7}, 2}, 6},
		{"3 num", args{[]int{5, 6, 7}, 3}, 5},
		{"3 num same", args{[]int{6, 6, 7}, 2}, 6},
		{"3 num same", args{[]int{6, 6, 7}, 3}, 6},
		{"4 num", args{[]int{5, 6, 7, 8}, 1}, 8},
		{"4 num", args{[]int{5, 6, 7, 8}, 2}, 7},
		{"4 num", args{[]int{5, 6, 7, 8}, 3}, 6},
		{"4 num", args{[]int{5, 6, 7, 8}, 4}, 5},
		{"4 num", args{[]int{7, 6, 7, 8}, 2}, 7},
		{"4 num", args{[]int{7, 6, 7, 8}, 3}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
