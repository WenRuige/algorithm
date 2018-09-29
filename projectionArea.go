package main

import (
	"fmt"
)

func projectionArea(grid [][]int) int {
	//bottom := len(grid)
	hang := 0
	hangmax := 0

	shu := 0
	shumax := 0
	botom := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > hang {
				hang = grid[i][j]
			}
			if grid[i][j] > 0 {
				botom = botom + 1
			}
		}
		hangmax = hang + hangmax
		hang = 0
	}

	// 按列遍历
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][i] > shu {
				shu = grid[j][i]
			}
		}
		shumax = shumax + shu
		shu = 0
	}

	fmt.Println(hangmax, shumax, botom)
	return hangmax + shumax + botom

}

func main() {

	grid := [][]int{{1, 0}, {0, 2}}
	fmt.Println(projectionArea(grid))
}
