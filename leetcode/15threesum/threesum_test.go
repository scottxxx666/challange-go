package _5threesum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"empty array", args{[]int{}}, [][]int{}},
		{"3 nums 1 solution", args{[]int{-1, 0, 1}}, [][]int{{-1, 0, 1}}},
		{"3 nums no solution", args{[]int{-1, 0, 2}}, [][]int{}},
		{"4 nums 1 solution", args{[]int{-1, 0, 1, 2}}, [][]int{{-1, 0, 1}}},
		{"4 nums 1 solution", args{[]int{-2, 0, 1, 2}}, [][]int{{-2, 0, 2}}},
		{"no duplicate", args{[]int{-2, -2, 0, 2}}, [][]int{{-2, 0, 2}}},
		{"not increase", args{[]int{2, -1, 0, -1}}, [][]int{{-1, -1, 2}}},
		{"not increase", args{[]int{-1, 0, 1, 2, -1, -4}},
			[][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
