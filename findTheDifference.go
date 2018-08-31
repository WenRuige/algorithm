package main

import "fmt"

func findTheDifference(s string, t string) byte {
	origin := make(map[string]int)
	res := uint8(0)
	for i := 0; i < len(s); i++ {
		_, ok := origin[string(s[i])]
		if ok {
			origin[string(s[i])] = origin[string(s[i])] + 1
		} else {
			origin[string(s[i])] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		_, ok := origin[string(t[i])]
		if ok {
			if origin[string(t[i])] > 1 {
				origin[string(t[i])] = origin[string(t[i])] - 1
			} else {
				delete(origin, string(t[i]))
			}
		} else {
			res = t[i]
		}
	}

	return res

}

func main() {
	res := findTheDifference("abc", "acbd")
	fmt.Println(res)
}
