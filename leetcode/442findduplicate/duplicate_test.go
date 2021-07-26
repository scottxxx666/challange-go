package _42findduplicate

import (
	"reflect"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1 num", args{[]int{1}}, []int{}},
		{"2 same num", args{[]int{1, 1}}, []int{1}},
		{"2 diff num", args{[]int{1, 2}}, []int{}},
		{"3 num", args{[]int{1, 2, 3}}, []int{}},
		{"3 num", args{[]int{1, 2, 1}}, []int{1}},
		{"3 num", args{[]int{2, 1, 1}}, []int{1}},
		{"4 num", args{[]int{1, 2, 3, 4}}, []int{}},
		{"4 num", args{[]int{1, 2, 1, 2}}, []int{1, 2}},
		{"4 num", args{[]int{2, 3, 4, 2}}, []int{2}},
		{"4 num", args{[]int{2, 2, 3, 4}}, []int{2}},
		{"4 num", args{[]int{1, 2, 3, 2}}, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
