package main

import (
	"fmt"
	"math"
)

func maxAreaOfIsland(grid [][]int) int {
	//dfs 深度优先搜索
	m := len(grid)
	n := len(grid[0])
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				max = int(math.Max(float64(max), float64(areaIsLand(grid, i, j))))
			}

		}
	}

	return max
}

func areaIsLand(grid [][]int, i, j int) int {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) && grid[i][j] != 0 {
		grid[i][j] = 0
		return 1 + areaIsLand(grid, i+1, j) + areaIsLand(grid, i-1, j) + areaIsLand(grid, i, j+1) + areaIsLand(grid, i, j-1)
	}
	return 0

}

func main() {
	nums := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(nums))
}
