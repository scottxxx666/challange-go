package _96nextgreater

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"no greater", args{[]int{1}, []int{1}}, []int{-1}},
		{"next greater", args{[]int{1}, []int{1, 2}}, []int{2}},
		{"next greater", args{[]int{1, 2, 3}, []int{1, 2, 3}}, []int{2, 3, -1}},
		{"not ordered", args{[]int{1, 3, 2, 4}, []int{4, 2, 1, 3}}, []int{3, -1, 3, -1}},
		{"not ordered", args{[]int{4, 1, 2}, []int{1, 3, 4, 2}}, []int{-1, 3, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
