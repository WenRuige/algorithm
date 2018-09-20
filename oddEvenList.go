package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	odd := new(ListNode)
	even := new(ListNode)
	headodd := odd
	headeven := even

	flag := 1
	for head != nil {
		temp := new(ListNode)
		temp.Val = head.Val
		if flag%2 == 0 {
			even.Next = temp
			even = even.Next
		} else {
			odd.Next = temp
			odd = odd.Next
		}
		head = head.Next
		flag++
	}

	odd.Next = headeven.Next
	return headodd.Next

}
