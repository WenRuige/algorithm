package main

import "fmt"

func transpose(A [][]int) [][]int {

	// 第一行的长度
	lens := len(A[0])
	// 声明一个二维数组来存储值
	num := [][]int{}
	for j := 0; j < lens; j++ {
		temp := []int{}
		for i := 0; i < len(A); i++ {
			temp = append(temp, A[i][j])
		}
		num = append(num, temp)
	}

	return num
}

func main() {

	nums := [][]int{{1, 3, 4}, {2, 6, 7}}
	res := transpose(nums)
	fmt.Println(res)
}
