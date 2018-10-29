package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	result := strings.Split(s, " ")
	resultLen := len(result)

	for i := 0; i < len(result); i++ {
		if len(result[i]) == 0 {
			resultLen--
		}
	}

	return resultLen
}

func main() {
	res := countSegments("          ")
	fmt.Println(res)
}
