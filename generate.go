package main

func generate(numRows int) [][]int {
	nums := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		nums[i] = make([]int, i+1)
		nums[i][0] = 1
		nums[i][i] = 1
	}

	for i := 2; i < len(nums); i++ {
		for j := 1; j < i; j++ {
			nums[i][j] = nums[i-1][j-1] + nums[i-1][j]
		}
	}

	return nums
}

func main() {
	generate(5)
}
