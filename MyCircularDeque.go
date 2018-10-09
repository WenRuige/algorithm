package main

type MyCircularDeque struct {
	Val []int
	Num int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func ConstructorQQ(k int) MyCircularDeque {
	deque := new(MyCircularDeque)
	deque.Num = k
	return *deque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	// 满了 return false
	if this.Num == len(this.Val) {
		return false
	}
	// 老的数组
	old := this.Val
	newArr := []int{value}
	newArr = append(newArr, old...)
	this.Val = newArr
	return true

}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.Num == len(this.Val) {
		return false
	}
	this.Val = append(this.Val, value)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if len(this.Val) == 0 {
		return false
	}
	this.Val = this.Val[1:]
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if len(this.Val) == 0 {
		return false
	}
	this.Val = this.Val[0 : len(this.Val)-1]
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if len(this.Val) == 0 {
		return -1
	}
	return this.Val[0]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if len(this.Val) == 0 {
		return -1
	}
	return this.Val[len(this.Val)-1]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if len(this.Val) == 0 {
		return true
	} else {
		return false
	}
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if len(this.Val) == this.Num {
		return true
	} else {
		return false
	}
}

func main() {
	obj := ConstructorQQ(4)
	obj.InsertFront(9)
	obj.DeleteLast()
	obj.GetRear()
	obj.GetFront()

}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
