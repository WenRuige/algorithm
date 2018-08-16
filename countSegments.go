package main

import "fmt"

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	total := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {

		}
	}
	return total + 1
}

func main() {
	res := countSegments("       ")
	fmt.Println(res)
}
