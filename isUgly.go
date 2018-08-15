package main

import "fmt"

func isUgly(num int) bool {

	if num <= 0 {
		return false
	}
	if num == 1 {
		return true
	}
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
			if num == 1 {
				return true
			}
		} else if num%3 == 0 {
			num = num / 3
			if num == 1 {
				return true
			}
		} else if num%5 == 0 {
			num = num / 5
			if num == 1 {
				return true
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	res := isUgly(-2147483648)
	fmt.Println(res)
}
