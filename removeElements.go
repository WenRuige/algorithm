package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}
