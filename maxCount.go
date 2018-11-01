package main

import (
	"fmt"
	"math"
)

func maxCount(m int, n int, ops [][]int) int {
	opsLen := len(ops)
	if opsLen < 1 {
		return m * n
	}
	mina := -math.MinInt32
	minb := -math.MinInt32

	for i := 0; i < len(ops); i++ {
		a := ops[i][0]
		b := ops[i][1]
		mina = int(math.Min(float64(a), float64(mina)))
		minb = int(math.Min(float64(b), float64(minb)))
	}
	return mina * minb
}

func main() {
	res := maxCount(3, 3, [][]int{{2, 2}})
	fmt.Println(res)
}
