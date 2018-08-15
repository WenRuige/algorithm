package main

import "fmt"

func rotateString(A string, B string) bool {

	if A == "" && B == "" {
		return true
	}
	newStr := ""
	for i := 0; i < len(A); i++ {
		newStr = A[i+1:] + A[:i+1]
		if newStr == B {
			return true
		}
	}

	return false
}

func main() {
	res := rotateString("", "")
	fmt.Println(res)
}
