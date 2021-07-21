package _52openthelock

import "testing"

func Test_openLock(t *testing.T) {
	type args struct {
		deadEnds []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"??", args{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openLock(tt.args.deadEnds, tt.args.target); got != tt.want {
				t.Errorf("openLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
