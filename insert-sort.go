package main

import (
	"fmt"

	"github.com/algorithm/common"
)

// 插入排序

/*

	对于一个数组A[0,n]的排序问题,可以认为A[0,n-1]的排序已经解决了

	最好的情况下需要n-1次操作
	最坏的情况下需要n(n-1)次操作



*/

// 插入排序时间复杂度:

func insertSort(array []int) {

	arrayLen := len(array)
	for i := 0; i < arrayLen; i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			common.Swap(array, j, j-1)
		}
	}

}

func main() {
	data := []int{4, 6, 3, 12, 65, 76, 5, 78, 9, 764, 323, 34, 43, 1}
	insertSort(data)

	fmt.Println(data)
}
