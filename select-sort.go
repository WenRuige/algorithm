package main

import (
	"fmt"

	"github.com/algorithm/common"
)

// 选择排序

// 找到数组中最小的元素,其次将它和数组第一个元素交换位置,再次,在剩下的元素中找出最小的元素并与第二位进行交换,如此往复直至将整个数组返回

// 时间复杂度:O(n^2)

func selectSort(arr []int) {
	dataStructLen := len(arr)

	for i := 0; i < dataStructLen; i++ {

		min := i // 最小元素的索引

		for j := i + 1; j < dataStructLen; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// 如果当前元素与最小元素索引不一致,表示有比当前元素更小的元素,需要交换
		if min != i {
			common.Swap(arr, min, i)
		}
	}

}

func main() {
	s := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	selectSort(s)
	fmt.Println(s)

}
