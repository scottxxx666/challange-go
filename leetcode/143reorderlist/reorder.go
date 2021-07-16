package _43reorderlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if slow.Next != nil {
		temp := slow.Next
		slow.Next = nil
		slow = temp
	}

	var prev *ListNode
	for slow != nil {
		n := slow.Next
		slow.Next = prev
		prev = slow
		slow = n
	}

	temp := head
	for temp != nil && prev != nil {
		n := temp.Next
		temp.Next = prev
		nn := prev.Next
		prev.Next = n
		temp = n
		prev = nn
	}

	for head != nil {
		fmt.Printf("%v", head)
		head = head.Next
	}
	println("\n---")
}
