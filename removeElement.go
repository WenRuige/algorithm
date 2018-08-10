package main

import "fmt"

func removeElement(nums []int, val int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {

		}
	}
	return 0
}

func main() {
	nums := []int{1, 2, 3, 4, 4, 5, 6, 7}
	res := removeElement(nums, 4)
	fmt.Println(res)
}
