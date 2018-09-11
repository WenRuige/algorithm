package main

import "fmt"

func compress(chars []byte) int {
	flag, total := 0, 0
	for i := 0; i < len(chars); i++ {
		if chars[i] == chars[i+1] {
			flag++
		} else {
			// 计算 times
			if flag > 0 {
				for flag > 0 {
					flag = flag / 10
				}
			} else {

			}

		}
	}

	return 0
}

func main() {
	nums := []byte{'a', 'v'}
	res := compress(nums)
	fmt.Println(res)
}
