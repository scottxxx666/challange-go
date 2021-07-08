package _40longestsubstringwithkchars

import "testing"

func Test_longestSubstring(t *testing.T) {
	type args struct {
		s      string
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 char solution 1", args{"a", 1}, 1},
		{"2 same char", args{"aa", 1}, 2},
		{"2 diff char target 1", args{"ab", 1}, 1},
		{"2 diff char target 2", args{"ab", 2}, 2},
		{"3 diff char target 2", args{"abc", 2}, 2},
		{"3 char 2 same", args{"aba", 2}, 3},
		{"3 char 3 same", args{"aaa", 1}, 3},
		{"1 char no solution", args{"a", 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubstring(tt.args.s, tt.args.target); got != tt.want {
				t.Errorf("longestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
