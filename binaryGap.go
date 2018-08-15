package main

func binaryGap(N int) int {
	sum := []int{}
	for N > 0 {
		pop := N % 2
		N = N / 2
		sum = append(sum, pop)
	}

	// 逆序
	total := []int{}
	for i := len(sum) - 1; i >= 0; i-- {
		total = append(total, sum[i])
	}

	// 计算1的数量
	position := []int{}
	for i := 0; i < len(total); i++ {
		if total[i] == 1 {
			position = append(position, i)
		}
	}

	if len(position) < 2 {
		return 0
	}
	max := -2 << 32
	for i := 0; i < len(position)-1; i++ {
		if position[i+1]-position[i] > max {
			max = position[i+1] - position[i]
		}
	}

	return max
}

func main() {
	binaryGap(6)
}
