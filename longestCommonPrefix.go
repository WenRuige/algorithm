package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	// 找出最短的,
	min := 1 << 31
	position := 0
	for k, v := range strs {
		if len(v) < min {
			min = len(v)
			position = k
		}
	}
	fmt.Println("start")

	endPosition := 0
	for i := 0; i < min; i++ {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] == strs[position][i] {

			} else {
				if i == 0 {
					endPosition = 0
					return ""

				} else {
					endPosition = i
					return strs[position][0:endPosition]
				}
			}
		}
	}
	return strs[position]
}

func main() {

	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}
