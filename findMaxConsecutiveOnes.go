package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	flag := 0
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 1 {
			if nums[i] == nums[i+1] {
				flag++
				if max < flag {
					max = flag
				}
			} else {
				flag = 0
			}
		}
	}
	//fmt.Println(max)
	return 0
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	res := findMaxConsecutiveOnes(nums)
	fmt.Println(res)
}
