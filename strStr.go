package main

import "fmt"

func strStr(haystack string, needle string) int {

	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	j := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[j] {
			fmt.Println(i, j)
			for k := i; k < len(haystack); k++ {
				if (haystack[k] == needle[j]) && j <= len(needle) {
					//fmt.Println(string(haystack[k]))
					j = j + 1
					//fmt.Printf("j's  num=%v|| value=%v||  len = %v || value =  %v  \n", j, string(haystack[k]), len(needle), string(needle[j]))
					if j == len(needle) {
						return i
					}
				} else {
					j = 0
					break
				}

			}
		}
	}

	return -1
}

func main() {
	position := strStr("mississippi", "issip")
	fmt.Printf("result=%v", position)
}
