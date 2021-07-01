package _6mergeIntervals

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"empty array", args{[][]int{}}, [][]int{}},
		{"1 element", args{[][]int{{0, 1}}}, [][]int{{0, 1}}},
		{"2 overlay element", args{[][]int{{0, 1}, {1, 2}}}, [][]int{{0, 2}}},
		{"2 overlay element", args{[][]int{{0, 2}, {1, 3}}}, [][]int{{0, 3}}},
		{"2 overlay element", args{[][]int{{0, 4}, {1, 3}}}, [][]int{{0, 4}}},
		{"2 overlay element", args{[][]int{{1, 2}, {0, 4}}}, [][]int{{0, 4}}},
		{"3 overlay element", args{[][]int{{1, 3}, {2, 4}, {4, 5}}}, [][]int{{1, 5}}},
		{"3 overlay element", args{[][]int{{4, 5}, {2, 4}, {1, 3}}}, [][]int{{1, 5}}},
		{"2 not overlay element", args{[][]int{{1, 2}, {3, 4}}}, [][]int{{1, 2}, {3, 4}}},
		{"3 not overlay element",
			args{[][]int{{5, 6}, {1, 2}, {3, 4}}},
			[][]int{{1, 2}, {3, 4}, {5, 6}}},
		{"merge to 2",
			args{[][]int{{2, 3}, {5, 7}, {1, 2}, {4, 6}}},
			[][]int{{1, 3}, {4, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
