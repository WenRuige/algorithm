package main

import "fmt"

func customSortString(S string, T string) string {
	str := ""

	newStr := make(map[string]int)

	for i := 0; i < len(T); i++ {
		_, ok := newStr[string(T[i])]
		if ok {
			newStr[string(T[i])] = newStr[string(T[i])] + 1
		} else {
			newStr[string(T[i])] = 1
		}
	}

	for i := 0; i < len(S); i++ {
		for j := 0; j < len(T); j++ {
			if S[i] == T[j] {
				str = str + string(S[i])
				newStr[string(S[i])] = newStr[string(S[i])] - 1
			}
		}
	}

	for k, v := range newStr {
		if v > 0 {
			for i := 0; i < v; i++ {
				str = str + k
			}
		}
	}
	return str
}

func main() {
	fmt.Println(customSortString("cba", "abbcdd"))
}
