package main

import "fmt"

// 在看看
func isIsomorphic(s string, t string) bool {
	num1 := []int{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] != s[j] {
				num1 = append(num1, 0)
			}
		}
	}

	fmt.Println(num1)

	return false
}

func main() {
	res := isIsomorphic("paper", "title")
	fmt.Println(res)
}
