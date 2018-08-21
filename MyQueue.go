package main

import "fmt"

type MyQueue struct {
	Val []int
}

/** Initialize your data structure here. */
func ConstructorI() MyQueue {
	res := new(MyQueue)
	return *res

}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Val = append(this.Val, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	last := this.Val[0]
	this.Val = this.Val[1:]
	return last

}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.Val[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {

	if len(this.Val) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	obj := ConstructorI()
	obj.Push(1)

	fmt.Println(obj.Pop())
	obj.Push(2)

	fmt.Println(obj.Peek())
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
