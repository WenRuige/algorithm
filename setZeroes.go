package main

import "fmt"

func setZeroes(matrix [][]int) {
	// 先循环一次,把为0的点都找到
	newArr := [][]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				newArr = append(newArr, []int{i, j})
			}
		}
	}

	// 计算出来m,n
	m := len(matrix[0])
	n := len(matrix)

	fmt.Println(newArr)

	for i := 0; i < len(newArr); i++ {
		for j := 0; j < m; j++ {
			matrix[newArr[i][0]][j] = 0
			//fmt.Println(newArr[i][0], j)
		}

		for q := 0; q < n; q++ {
			//fmt.Println(q, newArr[i][0])
			matrix[q][newArr[i][1]] = 0

		}
	}

}

func main() {
	nums := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(nums)
}
