package _67twosum2

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"2 number", args{[]int{1, 2}, 3}, []int{1, 2}},
		{"3 number(1, 2)", args{[]int{1, 2, 3}, 3}, []int{1, 2}},
		{"3 number(1, 3)", args{[]int{1, 2, 3}, 4}, []int{1, 3}},
		{"3 number(2, 3)", args{[]int{1, 2, 3}, 5}, []int{2, 3}},
		{"4 number(1, 3)", args{[]int{1, 2, 3, 4}, 4}, []int{1, 3}},
		{"4 number(2, 3)", args{[]int{0, 2, 3, 4}, 5}, []int{2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
