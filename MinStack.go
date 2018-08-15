package main

import "fmt"

type MinStack struct {
	Val []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	res := new(MinStack)
	return *res
}

func (this *MinStack) Push(x int) {
	this.Val = append(this.Val, x)
}

func (this *MinStack) Pop() {
	this.Val = this.Val[:len(this.Val)-1]
	//fmt.Println(this.Val)
}

func (this *MinStack) Top() int {
	return this.Val[len(this.Val)-1]
}

func (this *MinStack) GetMin() int {
	max := 1 << 32
	for i := 0; i < len(this.Val); i++ {
		if this.Val[i] < max {
			max = this.Val[i]
		}
	}
	return max

}

func main() {

	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	obj.Pop()
	res1 := obj.Top()
	fmt.Println(res1)

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
