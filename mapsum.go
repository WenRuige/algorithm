package main

import (
	"fmt"
	"strings"
)

type MapSum struct {
	Val map[string]int
}

/** Initialize your data structure here. */
func Constructorss() MapSum {
	res := new(MapSum)
	res.Val = make(map[string]int)
	return *res
}

func (this *MapSum) Insert(key string, val int) {
	this.Val[key] = val

}

func (this *MapSum) Sum(prefix string) int {
	total := 0
	for k, v := range this.Val {
		flag := strings.Contains(k, prefix)
		if flag {
			total = total + v
		}
	}
	return total
}

func main() {
	obj := Constructorss()
	obj.Insert("apple", 2)

	obj.Insert("aab", 33)
	obj.Insert("aewfwaefjeoawefjwoeajfowajfoewajfoawefjeowajfowaj", 111)
	obj.Insert("bb", 143)
	obj.Insert("cc", 144)
	obj.Insert("aab", 33)
	param := obj.Sum("aab")
	param2 := obj.Sum("ab")

	fmt.Println(param, param2)
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
