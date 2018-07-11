package main

import "fmt"
import "math"

// 二分查找

// 递归&非递归版本

// 思想:有序的数组,找到其中一个值

// 时间复杂度

// 入参: 有序数组  & 要查找的值
// 出参:

type DataStruct struct {
	Data []int
}

func (data *DataStruct) binaryChopSearch(key int) int {

	left, right, mid := 0, len(data.Data), 0
	for {
		mid = int(math.Floor(float64((left + right) / 2)))

		if key > data.Data[mid] {
			left = mid + 1
		} else if key < data.Data[mid] {
			right = mid - 1
		} else {
			break
		}
		// 如果左面大于右面,表示不存在
		if left > right {
			return -1
		}
	}
	return mid
}

func main() {

	res := DataStruct{[]int{1, 2, 3, 4, 5, 6, 7, 8}}
	final := res.binaryChopSearch(2)
	fmt.Println(final)

}
