package _77squareofasortedarray

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty array", args{[]int{}}, []int{}},
		{"1 positive", args{[]int{2}}, []int{4}},
		{"1 negative", args{[]int{-1}}, []int{1}},
		{"2 positive", args{[]int{2, 3}}, []int{4, 9}},
		{"1 negative 1 positive", args{[]int{-1, 2}}, []int{1, 4}},
		{"2 negative", args{[]int{-2, -1}}, []int{1, 4}},
		{"2 negative 1 positive", args{[]int{-2, -1, 1}}, []int{1, 1, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
