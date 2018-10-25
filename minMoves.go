package main

import (
	"fmt"
	"sort"
)

// 反向思考,减到0需要几步

/*

	1 2 3
    0 2 2
    0 1 1
    0 0 0


*/
func minMoves(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for _, v := range nums {
		sum = sum + v - nums[0]
	}
	return sum
}

func main() {

	fmt.Println(minMoves([]int{1, 2, 3}))
}
