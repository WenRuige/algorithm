package main

import "fmt"
import "math"

// 二分查找

// 递归&非递归版本

// 思想:有序的数组,找到其中一个值

// 时间复杂度: O(log2n)

// 时间复杂度分析:
/*
	对于N个元素的情况,
    1次2分   n/2
    2次2分   n/4(2*2)
    m次2分   n/(2^m)
    最坏的情况下排除到最后一个才产生结果   1 =  n/2(^m)     n = 2^m   所以时间复杂度为:O(log2n)

*/

// 入参: 有序数组  & 要查找的值
// 出参: 所在数组的位置,若不存在返回 -1

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

// 递归版本实现

func main() {

	res := DataStruct{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	final := res.binaryChopSearch(2)
	fmt.Println(final)

}
