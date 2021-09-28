package _22sortarraybyparity2

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"answer", args{[]int{2, 3}}, []int{2, 3}},
		{"2 nums exchange", args{[]int{3, 2}}, []int{2, 3}},
		{"4 nums", args{[]int{2, 3, 4, 5}}, []int{2, 3, 4, 5}},
		{"4 nums", args{[]int{3, 2, 4, 5}}, []int{2, 3, 4, 5}},
		{"4 nums", args{[]int{3, 2, 5, 4}}, []int{2, 3, 4, 5}},
		{"4 nums", args{[]int{4, 2, 5, 3}}, []int{4, 5, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParityII(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParityII() = %v, want %v", got, tt.want)
			}
		})
	}
}
