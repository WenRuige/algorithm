package main

//
//import "fmt"
//
//type MyHashSet struct {
//	Val []int
//}
//
///** Initialize your data structure here. */
//func Constructor() MyHashSet {
//	obj := new(MyHashSet)
//	return *obj
//}
//
//func (this *MyHashSet) Add(key int) {
//	this.Val = append(this.Val, key)
//}
//
//func (this *MyHashSet) Remove(key int) {
//	newArr := []int{}
//	for _, v := range this.Val {
//		if v == key {
//		} else {
//			newArr = append(newArr, v)
//		}
//	}
//
//	this.Val = newArr
//}
//
///** Returns true if this set did not already contain the specified element */
//func (this *MyHashSet) Contains(key int) bool {
//	for _, v := range this.Val {
//		if v == key {
//			return true
//		}
//	}
//	return false
//}
//
//func main() {
//	//Your MyHashSet object will be instantiated and called as such:
//	obj := Constructor()
//	obj.Add(1)
//	obj.Add(2)
//	obj.Contains(1)
//	obj.Contains(3)
//	obj.Add(2)
//	obj.Contains(2)
//	obj.Remove(2)
//	obj.Contains(2)
//
//	//obj.Remove(2)
//	param_3 := obj.Contains(2)
//	fmt.Println(param_3)
//}
//
///**
// * Your MyHashSet object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Add(key);
// * obj.Remove(key);
// * param_3 := obj.Contains(key);
// */
