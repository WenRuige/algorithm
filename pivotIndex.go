package main

import "fmt"

func pivotIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// 左半部分
		left := 0
		right := 0
		for j := 0; j < i; j++ {
			left = left + nums[j]
		}

		for k := i + 1; k < len(nums); k++ {
			right = right + nums[k]
		}

		fmt.Println(left, right)

		if left == right {
			return i
		}
	}

	return -1

}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	res := pivotIndex(nums)
	fmt.Println(res)
}
