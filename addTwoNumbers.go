package main

//
//import (
//	"math"
//)
//
///**
// * Definition for singly-linked list.
// * type ListNode struct {
// *     Val int
// *     Next *ListNode
// * }
// */
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	val1 := 0
//	cal1 := 1
//	val1 = l1.Val
//	point1 := l1.Next
//	for nil != point1 {
//		temp1 := point1.Val
//		val1 = val1 + temp1*int(math.Pow10(cal1))
//		cal1++
//		point1 = point1.Next
//	}
//
//	val2 := 0
//	cal2 := 1
//	val2 = l2.Val
//	point2 := l2.Next
//	for nil != point2 {
//		temp2 := point2.Val
//		val2 = val2 + temp2*int(math.Pow10(cal2))
//		cal2++
//		point2 = point2.Next
//	}
//
//	// 输出结果
//	res := val1 + val2
//
//	listNodeNew := new(ListNode)
//	head := listNodeNew
//	listNodeNew.Val = res % 10
//	res = res / 10
//	for res > 0 {
//		pop := res % 10
//		res = res / 10
//		small := new(ListNode)
//		small.Val = pop
//		listNodeNew.Next = small
//		listNodeNew = listNodeNew.Next
//	}
//
//	// 拆分
//
//	return head
//}
//
//func main() {
//
//	l1 := new(ListNode)
//	l1.Next = new(ListNode)
//	l1.Next.Next = new(ListNode)
//	l1.Val = 2
//	l1.Next.Val = 4
//	l1.Next.Next.Val = 3
//
//	l2 := new(ListNode)
//	l2.Next = new(ListNode)
//	l2.Next.Next = new(ListNode)
//	l2.Val = 5
//	l2.Next.Val = 6
//	l2.Next.Next.Val = 4
//
//	addTwoNumbers(l1, l2)
//}
