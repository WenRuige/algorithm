package main

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	max := 10000 / len(A)
	for i := 1; i < max; i++ {
		if strings.Contains(A, B) {
			return i
		}
		A = A + A

	}
	return 0
}

/*
"abc"
"cabcabca"

*/
func main() {
	res := repeatedStringMatch("abcd", "cdabcdab")
	fmt.Println(res)
}
