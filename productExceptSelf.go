package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	output := make([]int, n)

	for mult, i := 1, 0; i < n; i++ {
		output[i] = mult
		mult *= nums[i]
	}
	fmt.Println(output)

	for mult, i := 1, n-1; i >= 0; i-- {
		output[i] *= mult
		mult *= nums[i]
	}

	return output
}

func main() {
	nums := []int{1, 0, 2}
	fmt.Println(productExceptSelf(nums))
}
