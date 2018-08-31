package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	//flag := false
	max := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			temp++
			if temp >= max {
				max = temp
			}
			//flag = true

		} else {
			//flag = false
			temp = 0
		}
	}
	return max
}

func main() {
	nums := []int{0}
	res := findMaxConsecutiveOnes(nums)
	fmt.Println(res)
}
