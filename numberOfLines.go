package main

import (
	"fmt"
)

func numberOfLines(widths []int, S string) []int {
	var returnArr []int
	lines := 0
	lineLength := 0
	if len(S) > 0 {
		lines = lines + 1
	}
	for _, char := range S {
		if lineLength+widths[char-97] > 100 {
			lines = lines + 1
			lineLength = 0
		}
		lineLength = lineLength + widths[char-97]
	}

	returnArr = append(returnArr, lines)
	returnArr = append(returnArr, lineLength)
	return returnArr
}

func main() {
	width := []int{7, 5, 4, 7, 10, 7, 9, 4, 8, 9, 6, 5, 4, 2, 3, 10, 9, 9, 3, 7, 5, 2, 9, 4, 8, 9}
	str := "zlrovckbgjqofmdzqetflraziyvkvcxzahzuuveypqxmjbwrjvmpdxjuhqyktuwjvmbeswxuznumazgxvitdrzxmqzhaaudztgie"
	res := numberOfLines(width, str)
	fmt.Println(res)
}
