package main

import "fmt"

type MyCircularQueue struct {
	Val []int
	Num int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func ConstructorQi(k int) MyCircularQueue {
	circular := new(MyCircularQueue)
	circular.Num = k
	return *circular
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.Val) == this.Num {
		return false
	} else {
		this.Val = append(this.Val, value)
		return true
	}
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if len(this.Val) == 0 {
		return false
	} else {
		this.Val = this.Val[1:]
		return true
	}
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if len(this.Val) == 0 {
		return -1
	} else {
		return this.Val[0]
	}

}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if len(this.Val) == 0 {
		return -1
	} else {
		return this.Val[len(this.Val)-1]
	}
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if len(this.Val) == 0 {
		return true
	}
	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if len(this.Val) == this.Num {
		return true
	}
	return false
}



func main(){
	obj:=ConstructorQi(8)
	obj.EnQueue(3)
	obj.EnQueue(9)
	obj.EnQueue(5)
	obj.EnQueue(0)
	obj.DeQueue()
	obj.DeQueue()
	fmt.Println(obj.Rear())

}
/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
