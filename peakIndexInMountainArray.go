package main

import "fmt"

func peakIndexInMountainArray(A []int) int {
	max := -1 << 32
	flag := 0
	for i := 0; i < len(A); i++ {
		if A[i] > max {

			max = A[i]
			flag = i
		}
	}

	//fmt.Println(max)
	return flag
}

func main() {
	nums := []int{0, 2, 1, 0}
	res := peakIndexInMountainArray(nums)
	fmt.Println(res)
}
