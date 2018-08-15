package main

//
///**
// * Definition for singly-linked list.
// * type ListNode struct {
// *     Val int
// *     Next *ListNode
// * }
// */
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	newListNode := new(ListNode)
//	head := newListNode
//
//	if l1 != nil && l2 == nil {
//		newListNode.Next = l1
//	}
//
//	if l1 == nil && l2 != nil {
//		newListNode.Next = l2
//	}
//
//	for l1 != nil && l2 != nil {
//		if l1.Val < l2.Val {
//			newListNode.Next = l1
//			l1 = l1.Next
//			newListNode = newListNode.Next
//		} else {
//			newListNode.Next = l2
//			l2 = l2.Next
//			newListNode = newListNode.Next
//		}
//	}
//
//	if l1 != nil {
//		newListNode.Next = l1
//	} else {
//		newListNode.Next = l2
//	}
//	return head.Next
//
//}
//
//func search(list *ListNode) {
//
//	for list != nil {
//		fmt.Println(list.Val)
//		list = list.Next
//	}
//}
//
//func main() {
//
//	l1 := new(ListNode)
//	//l1.Next = new(ListNode)
//	//l1.Next.Next = new(ListNode)
//	l1.Val = 1
//	//l1.Next.Val = 2
//	//l1.Next.Next.Val = 4
//
//	l2 := new(ListNode)
//	//l2.Next = new(ListNode)
//	//l2.Next.Next = new(ListNode)
//	//l2.Next.Next.Next = new(ListNode)
//	l2.Val = 2
//	//l2.Next.Val = 3
//	//l2.Next.Next.Val = 4
//	//l2.Next.Next.Next.Val = 5
//	res := mergeTwoLists(l1, l2)
//	search(res)
//
//}
