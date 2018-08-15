package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	// 如果小于原数组,返回
	if len(nums)*len(nums[0]) < r*c {
		return nums
	}

}

func main() {
	nums := [][]int{{1, 2}, {3, 4}}
	res := matrixReshape(nums, 1, 4)
	fmt.Println(res)
}
