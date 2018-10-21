package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func middleNode(head *ListNode) *ListNode {
	// 求这个链表的长
	tmp := head
	length := 0
	for tmp != nil {
		tmp = tmp.Next
		length ++
	}
	// 中间节点,
	middle := length / 2
	newLength := 0
	for head != nil {
		if middle == newLength {
			break
		}
		head = head.Next
		newLength ++

	}
	return head

}

func main() {
	l := new(ListNode)
	l.Val = 1
	l.Next = new(ListNode)
	l.Next.Val = 2
	l.Next.Next = new(ListNode)
	l.Next.Next.Val = 3
	l.Next.Next.Next = new(ListNode)
	l.Next.Next.Next.Val = 4
	l.Next.Next.Next.Next = new(ListNode)
	l.Next.Next.Next.Next.Val = 5
	res := middleNode(l)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
