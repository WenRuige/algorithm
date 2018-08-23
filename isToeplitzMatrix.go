package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {

	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}

	return true
}

func main() {

	nums := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	res := isToeplitzMatrix(nums)
	fmt.Println(res)
}
