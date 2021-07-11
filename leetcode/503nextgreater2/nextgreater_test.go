package _03nextgreater2

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1 num", args{[]int{1}}, []int{-1}},
		{"2 num", args{[]int{1, 2}}, []int{2, -1}},
		{"asc", args{[]int{1, 2, 3}}, []int{2, 3, -1}},
		{"no order", args{[]int{1, 3, 2, 0, 4}}, []int{3, 4, 4, 4, -1}},
		{"circular", args{[]int{1, 3, 2, 5, 4}}, []int{3, 5, 5, -1, 5}},
		{"circular", args{[]int{1, 5, 2, 3, 4}}, []int{5, -1, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
