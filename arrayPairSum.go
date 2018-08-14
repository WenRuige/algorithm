package main

import "fmt"

func arrayPairSum(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	total := 0
	for i := 0; i < len(nums); i++ {
		total = total + nums[i]
		i++
	}

	return total

}

func main() {
	nums := []int{1, 2}
	res := arrayPairSum(nums)
	fmt.Println(res)
}
