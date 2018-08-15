package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	cur := new(ListNode)
	h := cur
	//cur.Next = nil
	for head != nil {
		// 临时创建的链表
		temp := new(ListNode)
		temp.Val = head.Val
		// 创建结束
		temp.Next = cur.Next
		cur.Next = temp
		head = head.Next
	}

	return h.Next
}

func search_1(list *ListNode) {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
func main() {
	l1 := new(ListNode)
	l1.Next = new(ListNode)
	l1.Next.Next = new(ListNode)
	l1.Val = 1
	l1.Next.Val = 2
	l1.Next.Next.Val = 3

	res := reverseList(l1)
	search_1(res)
	//fmt.Println(s)
}
