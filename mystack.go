package main

type MyStack struct {
	Val []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	res := new(MyStack)
	return *res
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Val = append(this.Val, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	temp := this.Val[len(this.Val)-1]
	this.Val = this.Val[:len(this.Val)-1]

	return temp
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Val[len(this.Val)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.Val) == 0 {
		return true
	}
	return false
}

func main() {
	obj := Constructor()
	obj.Push(1)

}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
