package main

import (
	"fmt"
)

func maximumProduct(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if float64(nums[j]) > float64(nums[j+1]) {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}

	if nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3] > nums[0]*nums[1]*nums[len(nums)-1] {
		return nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	} else {
		return nums[0] * nums[1] * nums[len(nums)-1]
	}

}

func main() {
	nums := []int{-4, -3, -2, -1, 60}
	res := maximumProduct(nums)
	fmt.Println(res)
}
