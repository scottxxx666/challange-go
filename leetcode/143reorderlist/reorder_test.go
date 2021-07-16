package _43reorderlist

import "testing"

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{&ListNode{Val: 1}}},
		{"2", args{&ListNode{Val: 1, Next: &ListNode{Val: 2}}}},
		{"3", args{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		{"4", args{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next:
		&ListNode{Val: 3, Next: &ListNode{Val: 4}}}}}},
		{"5", args{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next:
		&ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}},
	}
	for _, tt := range tests {
		reorderList(tt.args.head)
	}
}
