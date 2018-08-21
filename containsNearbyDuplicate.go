package main

import (
	"fmt"
	"math"
)

func containsNearbyDuplicate(nums []int, k int) bool {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				//fmt.Println(1)
				res := i - j
				if int(math.Abs(float64(res))) <= k {
					return true
				}
			}
		}
	}

	return false

}

func main() {

	nums := []int{99, 99}

	res := containsNearbyDuplicate(nums, 2)
	fmt.Println(res)
}
