package _76middleoflinkedlist

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1 node", args{
			&ListNode{Val: 1},
		}, &ListNode{Val: 1}},
		{"2 nodes", args{
			&ListNode{Val: 1, Next: &ListNode{Val: 2}},
		}, &ListNode{Val: 2}},
		{"3 nodes", args{
			&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		}, &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		{"4 nodes", args{
			&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		}, &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{"5 nodes", args{
			&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
		}, &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
