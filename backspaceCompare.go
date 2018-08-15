package main

import "fmt"

func backspaceCompare(S string, T string) bool {
	a := []byte{}
	for i := 0; i < len(S); i++ {
		if string(S[i]) == "#" {

			if len(a) <= 0 {

			} else {
				a = a[:len(a)-1]
			}
		} else {
			a = append(a, S[i])
		}
	}

	b := []byte{}
	for i := 0; i < len(T); i++ {
		if string(T[i]) == "#" {
			//fmt.Println()
			if len(b) <= 0 {

			} else {
				b = b[:len(b)-1]
			}
		} else {
			b = append(b, T[i])
		}
	}

	if string(a) == string(b) {
		return true
	}
	return false
}

func main() {
	res := backspaceCompare("xywrrmp", "xywrrmu#p")
	fmt.Println(res)
}
