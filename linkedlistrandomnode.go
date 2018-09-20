package main

import "time"

package main

import (
"fmt"
"math/rand"
"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	Val []int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	solution := new(Solution)
	for head != nil {
		solution.Val = append(solution.Val, head.Val)
		head = head.Next
	}
	return *solution
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {

	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	return this.Val[r.Intn(len(this.Val))]

}

func main() {

	node := new(ListNode)
	node.Val = 1
	obj := Constructor(node)
	fmt.Println(obj.GetRandom())
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
