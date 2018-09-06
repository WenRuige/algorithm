package main

import (
	"fmt"
)

func isHappy(n int) bool {
	newArr := make(map[int]int)
	for n != 1 {
		res := getNextHappyNum(n)
		n = res
		fmt.Println(res)
		_, ok := newArr[res]
		if ok {
			return false
		} else {
			newArr[res] = 1
		}
	}
	return true
}

func getNextHappyNum(n int) int {
	sum := 0
	for n > 0 {
		pop := n % 10
		sum = sum + pop*pop
		n = n / 10
	}
	return sum
}

func main() {
	res := isHappy(19)
	fmt.Println(res)
}
