package main

import (
	"fmt"
	"math"
)

func shortestToChar(S string, C byte) []int {
	data := []int{}
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			data = append(data, i)
		}
	}
	res := []int{}
	for i := 0; i < len(S); i++ {
		flag := 2 << 32
		for j := 0; j < len(data); j++ {
			// 位置
			value := int(math.Abs(float64(data[j] - i)))
			if value < flag {
				flag = value
			}
		}
		res = append(res, flag)

	}

	return res
}

func main() {
	res := shortestToChar("loveleetcode", 'e')
	fmt.Println(res)
}
