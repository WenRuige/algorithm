package main

// 杨辉三角
func getRow(rowIndex int) []int {
	//构造外围
	result := [][]int{{1}, {1, 1}}
	for i := 2; i <= rowIndex; i++ {
		tmp := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				tmp = append(tmp, 1)
			} else {
				//计算上一行的树
				res := result[i-1][j] + result[i-1][j-1]
				tmp = append(tmp, res)
			}
		}
		result = append(result, tmp)
		tmp = []int{}
	}

	return result[rowIndex]
}

func main() {
	getRow(3)
}
