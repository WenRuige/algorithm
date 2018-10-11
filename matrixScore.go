package main

import (
	"fmt"
	"math"
)

/*
	判断逻辑
	1.第一列必须为1
	2.当第一列为1之后剩下的列尽可能保证过多的1


*/
func matrixScore(A [][]int) int {
	// 行转换,确认第一列都是1
	for i := 0; i < len(A); i++ {
		if A[i][0] == 0 {
			for j := 0; j < len(A[i]); j++ {
				if A[i][j] == 0 {
					A[i][j] = 1
				} else {
					A[i][j] = 0
				}
			}
		}
	}

	// 对每一列进行遍历,排除第一列
	for i := 1; i < len(A[0]); i++ {
		countZero := 0
		countOne := 0
		for j := 0; j < len(A); j++ {
			if A[j][i] == 0 {
				countZero++
			} else {
				countOne++
			}
		}
		// 如果0的个数大于1的个数,遍历当前列
		if countZero > countOne {
			for j := 0; j < len(A); j++ {
				if A[j][i] == 0 {
					A[j][i] = 1
				} else {
					A[j][i] = 0
				}
			}
		}

	}
	fmt.Println(A)
	// 对于每列进行二进制运算
	sum := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			sum = A[i][j]*int(math.Pow(2, float64(len(A[i])-j-1))) + sum
		}
	}

	return sum
}

func main() {
	//nums := [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}
	nums := [][]int{{0, 1}, {1, 1}}
	fmt.Println(matrixScore(nums))
}
