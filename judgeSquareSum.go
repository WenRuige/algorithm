package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	res := make(map[int]int)
	for i := 0; i <= int(math.Ceil(math.Sqrt(float64(c)))); i++ {
		res[i*i] = 1
		_, ok := res[c-i*i]
		if ok {
			return true
		}
	}
	return false
}

func main() {
	res := judgeSquareSum(5)
	fmt.Println(res)
}
