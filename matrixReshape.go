package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {
	// 如果小于原数组,返回
	if len(nums)*len(nums[0]) < r*c {
		return nums
	}
	// 分配一个数组来存储
	// golang 声明一个二维数组

	newArr := []int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			newArr = append(newArr, nums[i][j])
		}
	}

	finalArr := [][]int{}
	tempArr := []int{}
	temp := 0
	for i := 0; i < len(newArr); i++ {

		if temp == c-1 {
			tempArr = append(tempArr, newArr[i])
			finalArr = append(finalArr, tempArr)
			tempArr = []int{}
			temp = 0
		} else {
			tempArr = append(tempArr, newArr[i])
			temp++
		}
	}

	return finalArr
}

func main() {
	nums := [][]int{{1, 2}, {3, 4}}
	res := matrixReshape(nums, 1, 4)
	fmt.Println(res)
}
