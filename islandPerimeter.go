package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	land := 0
	neighbor := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				land++

				if i+1 < len(grid) && grid[i+1][j] == 1 {
					neighbor++
				}
				if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
					neighbor++
				}

				// 找一下左面/右面/上面/下面
			}
		}
	}
	return land*4 - neighbor*2
}

func main() {
	nums := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {1, 1, 1, 0}, {1, 1, 0, 0}}
	res := islandPerimeter(nums)
	fmt.Println(res)
}
