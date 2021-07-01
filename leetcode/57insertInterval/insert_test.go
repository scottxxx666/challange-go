package _7insertInterval

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"empty", args{[][]int{}, []int{}}, [][]int{{}}},
		{"no merge", args{[][]int{{1, 2}, {3, 4}}, []int{5, 6}},
			[][]int{{1, 2}, {3, 4}, {5, 6}}},
		{"merge 1", args{[][]int{{1, 2}, {5, 6}}, []int{2, 3}},
			[][]int{{1, 3}, {5, 6}}},
		{"merge 1 (include)", args{[][]int{{1, 2}, {5, 6}}, []int{0, 3}},
			[][]int{{0, 3}, {5, 6}}},
		{"merge 2", args{[][]int{{1, 2}, {5, 6}, {7, 8}}, []int{2, 5}},
			[][]int{{1, 6}, {7, 8}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
