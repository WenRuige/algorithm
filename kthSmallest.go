package main

import (
	"fmt"
	"sort"
)

func kthSmallest(matrix [][]int, k int) int {
	newArr := []int{}
	for i := 0; i < len(matrix); i++ {
		newArr = append(newArr, matrix[i]...)
	}
	sort.Ints(newArr)
	return newArr[k-1]
}

func main() {
	nums := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	res := kthSmallest(nums, 8)
	fmt.Println(res)
}
