package main

import "fmt"

func findLUSlength(a string, b string) int {
	aLength := len(a)
	bLength := len(b)

	if a == b {
		return -1
	}
	if aLength > bLength {
		return aLength
	} else {
		return bLength
	}

}

func main() {
	res := findLUSlength("absdasd", "asdasdasdsd")
	fmt.Println(res)
}
