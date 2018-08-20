package main

import "testing"

func Test_hourglassSum(t *testing.T) {
	type args struct {
		arr [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{[][]int32{
				[]int32{int32(1), int32(1), int32(1), int32(1), int32(1), int32(1)},
				[]int32{int32(1), int32(1), int32(1), int32(1), int32(1), int32(1)},
				[]int32{int32(1), int32(1), int32(1), int32(1), int32(1), int32(1)},
				[]int32{int32(1), int32(1), int32(1), int32(1), int32(1), int32(1)},
				[]int32{int32(1), int32(1), int32(1), int32(1), int32(1), int32(1)},
				[]int32{int32(1), int32(1), int32(1), int32(1), int32(1), int32(1)},
			}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hourglassSum(tt.args.arr); got != tt.want {
				t.Errorf("hourglassSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
