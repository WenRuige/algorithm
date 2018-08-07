package main

import (
	"fmt"
)

//import "strings"

func romanToInt(s string) int {
	//I, V, X, L, C, D, M := 1, 5, 10, 50, 100, 500, 1000
	//strings.Split("")
	// IV,IX :=4,9
	// XL,XC:=40,90
	// CD,CM:=400,900
	total := 0
	temp := []rune(s)
	for i := 0; i < len(s)-1; i++ {
		res := string(s[i]) + string(s[i+1])
		if res == "IV" {
			total = total + 4
			temp[i] = 0
			temp[i+1] = 0
		}
		if res == "IX" {
			total = total + 9
			temp[i] = 0
			temp[i+1] = 0
		}
		if res == "XL" {
			total = total + 40
			temp[i] = 0
			temp[i+1] = 0
		}
		if res == "XC" {
			total = total + 90
			temp[i] = 0
			temp[i+1] = 0
		}
		if res == "CD" {
			total = total + 400
			temp[i] = 0
			temp[i+1] = 0
		}
		if res == "CM" {
			total = total + 900
			temp[i] = 0
			temp[i+1] = 0
		}

	}

	for _, v := range temp {

		if string(v) == "I" {
			total = total + 1
		}
		if string(v) == "V" {
			total = total + 5
		}
		if string(v) == "X" {
			total = total + 10
		}
		if string(v) == "L" {
			total = total + 50
		}
		if string(v) == "C" {
			total = total + 100
		}
		if string(v) == "D" {
			total = total + 500
		}
		if string(v) == "M" {
			total = total + 1000
		}
	}

	return total

}

func main() {
	res := romanToInt("IXIIIIII")
	fmt.Println(res)
}
