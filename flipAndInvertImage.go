package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {

	nums := [][]int{}
	for i := 0; i < len(A); i++ {
		temp := []int{}
		for j := len(A[i]) - 1; j >= 0; j-- {
			temp = append(temp, A[i][j])
		}
		// 反过来
		//fmt.Println(temp)
		for k := 0; k < len(temp); k++ {
			if temp[k] == 0 {
				temp[k] = 1
			} else if temp[k] == 1 {
				temp[k] = 0
			}

		}
		nums = append(nums, temp)
	}

	return nums

}

func main() {
	num := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	res := flipAndInvertImage(num)
	fmt.Println(res)
}
