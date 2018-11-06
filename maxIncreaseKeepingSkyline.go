package main

import (
	"fmt"
)

func maxIncreaseKeepingSkyline(grid [][]int) int {
	// 水平天际线
	Vertical, Level := []int{}, []int{}
	for i := 0; i < len(grid); i++ {
		max := -1
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > max {
				max = grid[i][j]
			}
		}
		Level = append(Level, max)
	}
	// 竖直天际线
	for i := 0; i < len(grid[0]); i++ {
		max := 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] > max {
				max = grid[j][i]
			}
		}
		Vertical = append(Vertical, max)
	}
	// 算结果
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// 这个值要小于水平又小于竖直
			if grid[i][j] <= Level[i] || grid[i][j] <= Vertical[i] {
				add := mins(Level[i], Vertical[j])
				result += add - grid[i][j]
			}

		}

	}
	return result
}
func mins(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	res := maxIncreaseKeepingSkyline([][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}})
	fmt.Println(res)
}
