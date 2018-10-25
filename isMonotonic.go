package main

import "fmt"

func isMonotonic(A []int) bool {
	// 是否单调递增
	isDiZeng := true
	// 是否递减
	isDiJian := true
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			isDiZeng = false
			break
		}
	}

	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			isDiJian = false
			break
		}
	}
	return isDiZeng || isDiJian
}

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
}
