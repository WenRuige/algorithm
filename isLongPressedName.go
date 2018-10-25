package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
	k := 0
	for i := 0; i < len(name); i++ {
		flag := false
		for j := k; j < len(typed); j++ {
			if string(name[i]) == string(typed[j]) {
				fmt.Println(string(name[i]), string(typed[j]))
				k++
				flag = true
			} else {
				if !flag {
					return false
				}
				break
			}

		}
	}
	return true
}

func main() {
	fmt.Println(isLongPressedName("leelee", "lleeelee"))
}
