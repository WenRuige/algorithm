package main

import (
	"fmt"
	"os"
)

type LRUCache struct {
	Val    map[int]int //
	Data   []int
	Length int
}

func ConstructorGG(capacity int) LRUCache {
	lrucache := new(LRUCache)
	lrucache.Val = make(map[int]int)
	lrucache.Length = capacity
	return *lrucache
}

func (this *LRUCache) Get(key int) int {

	// 需要进行顺序调整
	flag := 0
	for i := 0; i < len(this.Val); i++ {
		if this.Val[i] == key {
			flag = i
		}
	}

	//fmt.Printf("------%v\n", flag)
	//fmt.Printf("++++++%v", this.Val[key])
	this.Val[0], this.Val[flag] = this.Val[flag], this.Val[0]

	return this.Val[key]
}

func (this *LRUCache) Put(key int, value int) {

	if len(this.Val) < this.Length {
		_, ok := this.Val[key]
		if !ok {
			this.Val[key] = value
			this.Data = append(this.Data, key)
		}
	} else {
		_, ok := this.Val[key]
		if !ok {
			lastData := this.Val[this.Length-1]
			delete(this.Val, lastData)
			this.Val[key] = value
			this.Data = append(this.Data, key)
		}
	}
}

func main() {

	lru := ConstructorGG(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // 返回  1

	os.Exit(1)
	lru.Put(3, 3)           // 该操作会使得密钥 2 作废
	fmt.Println(lru.Get(2)) // 返回 -1 (未找到)
	lru.Put(4, 4)           // 该操作会使得密钥 1 作废
	fmt.Println(lru.Get(1)) // 返回 -1 (未找到)
	fmt.Println(lru.Get(3)) // 返回  3
	fmt.Println(lru.Get(4)) // 返回  4
	fmt.Println(lru.Get(1))

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
