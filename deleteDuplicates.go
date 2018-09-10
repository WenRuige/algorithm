package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func deleteDuplicates(head *ListNode) *ListNode {

	if nil == head || head.Next == nil {
		return head
	}
	prePtr := head
	curPtr := head.Next
	lastValue := prePtr.Val
	for curPtr != nil {
		if curPtr.Val != lastValue {
			lastValue = curPtr.Val
			prePtr = curPtr
			curPtr = curPtr.Next

		} else {
			prePtr.Next = curPtr.Next
			curPtr = curPtr.Next
		}
	}
	return head
}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 1
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 2

	res := deleteDuplicates(l1)
	fmt.Println(res)

}
