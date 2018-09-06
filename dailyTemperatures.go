package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {

	stack := []int{}
	res := []int{}
	for i := len(temperatures) - 1; i >= 0; i-- {
		stack = append(stack, temperatures[i])
		temp := temperatures[i]
		flag := 0
		flagTemp := false
		for j := len(stack) - 1; j >= 0; j-- {
			if stack[j] > temp {
				flagTemp = true
				res = append(res, flag)
				break
			} else {
				flag++
			}
		}
		if !flagTemp {
			res = append(res, 0)
			flagTemp = false
		}
	}
	newRes := []int{}
	for i := len(res) - 1; i >= 0; i-- {
		newRes = append(newRes, res[i])
	}

	return newRes
}

func main() {
	nums := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(nums)
	fmt.Println(res)
}
