package main

import "fmt"

//func maxSubArray(nums []int) int {
//	max := -1000000000
//	for i := 0; i < len(nums); i++ {
//		for j := i; j < len(nums); j++ {
//			s := 0
//			for k := i; k <= j; k++ {
//				s = s + nums[k]
//			}
//			if s > max {
//				max = s
//			}
//		}
//	}
//	return max
//}

func maxSubArray(nums []int) int {
	max := -2 << 32
	for i := 0; i < len(nums); i++ {
		thisSum := 0
		for j := i; j < len(nums); j++ {
			thisSum = thisSum + nums[j]

			if max < thisSum {
				max = thisSum
			}
		}
	}
	return max
}

func main() {
	data := []int{1, 2, 3, -1}
	res := maxSubArray(data)
	fmt.Println(res)
}
