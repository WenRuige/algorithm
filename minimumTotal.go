package main

import (
	"fmt"
	"sort"
)

func minimumTotal(triangle [][]int) int {
	total := 0
	for i := 0; i < len(triangle); i++ {
		sort.Ints(triangle[i])
		fmt.Println(triangle[i])
		total = total + triangle[i][0]
	}
	return total

}

func main() {

	//nums := [][]int{{1, 3, 2}, {4, 6}, {1142, 123}}
	nums := [][]int{{-1}, {2, 3}, {-1, -1, -3, 4}}
	fmt.Println(minimumTotal(nums))
}
