package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	var dfs func(int, []int)

	dfs = func(idx int, temp []int) {
		t := make([]int, len(temp))
		copy(t, temp)
		// 没有以上两行，答案就是错的
		// 因为temp的底层数组在递归过程中，不停地修改
		// 程序结束时，temp的底层数组的值，全部是 nums 的最大值。
		res = append(res, t)
		//fmt.Println(t, idx)

		for i := idx; i < len(nums); i++ {
			fmt.Println(i)
			if i == idx || nums[i] != nums[i-1] {
				dfs(i+1, append(temp, nums[i]))
			}
		}
	}

	temp := make([]int, 0, len(nums))
	dfs(0, temp)

	return res
}

func main() {
	nums := []int{1, 2, 3}
	res := subsetsWithDup(nums)
	fmt.Println(res)
}
