package main

import (
	"fmt"
	"sort"
)

// 前k个高频元素

type num struct {
	num int
	fre int
}

type NumArr []*num

func (n NumArr) Len() int {
	return len(n)
}

func (n NumArr) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n NumArr) Less(i, j int) bool {
	if n[i].fre > n[j].fre {
		return true
	}
	return false
}

func topKFrequent3(nums []int, k int) []int {
	dataMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := dataMap[nums[i]]
		if ok {
			dataMap[nums[i]] = dataMap[nums[i]] + 1
		} else {
			dataMap[nums[i]] = 1
		}
	}
	fw := make(NumArr, 0, len(dataMap))
	for k, v := range dataMap {
		fw = append(fw, &num{k, v})
	}
	sort.Sort(fw)
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, fw[i].num)
	}

	return res
}

func main() {

	res := topKFrequent3([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Println(res)
}
