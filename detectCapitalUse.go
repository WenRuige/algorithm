package main

import "fmt"

func detectCapitalUse(word string) bool {
	total := 0
	for i := 0; i < len(word); i++ {
		if word[i] >= 65 && word[i] <= 92 {
			total++
		}
	}
	// 如果全是大写
	if len(word) == total {
		return true
	}
	// 首字母
	if total == 1 && word[0] >= 65 && word[0] <= 92 {
		return true
	}

	//全是小写

	if total == 0 {
		return true
	}

	return false
}

func main() {
	res := detectCapitalUse("Flag")
	fmt.Println(res)
}
