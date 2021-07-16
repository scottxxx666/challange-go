package _57circulararrayloop

import "testing"

func Test_circularArrayLoop(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1 num", args{[]int{1}}, false},
		{"2 num loop", args{[]int{1, 1}}, true},
		{"2 num no loop (all pos/neg)", args{[]int{1, -1}}, false},
		{"2 num circular loop", args{[]int{1, 3}}, true},
		{"2 num loop length should > 1", args{[]int{1, 2}}, false},
		{"2 num no loop", args{[]int{-1, 1}}, false},
		{"3 num loop", args{[]int{1, 2, 3}}, true},
		{"4 num no loop", args{[]int{1, 2, 3, 4}}, false},
		{"4 num loop", args{[]int{1, 2, 3, 5}}, true},
		{"5 num loop", args{[]int{1, 2, 2, -1, 3}}, true},
		{"5 num no loop", args{[]int{1, 2, 3, -1, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circularArrayLoop(tt.args.nums); got != tt.want {
				t.Errorf("circularArrayLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
