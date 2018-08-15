package main

import "fmt"

type MyHashMap struct {
	Val map[int]int
}

/** Initialize your data structure here. */
func Constructors() MyHashMap {
	res := new(MyHashMap)
	res.Val = make(map[int]int)
	return *res
	//return *new(MyHashMap)
}

/** value will always be positive. */
func (this *MyHashMap) Put(key int, value int) {
	this.Val[key] = value

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	_, ok := this.Val[key]
	if !ok {
		return -1
	}
	return this.Val[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	delete(this.Val, key)
}

func main() {
	obj := Constructors()
	obj.Put(1, 2)
	fmt.Println(obj.Get(4))
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
