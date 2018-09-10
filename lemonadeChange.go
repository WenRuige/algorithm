package main

import "fmt"

func lemonadeChange(bills []int) bool {

	count5, count10, count20 := 0, 0, 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			count5++
		} else if bills[i] == 10 {
			count10++
			if count5 > 0 {
				count5--
			} else {
				return false
			}
		} else {
			count20++
			if count10 > 0 {
				count10--
				if count5 > 0 {
					count5--
				} else {
					return false
				}
			} else {

				if count5 > 3 {
					count5 = count5 - 3
				} else {
					return false
				}

			}
			//20的逻辑
		}
	}
	return true
}

func main() {

	bills := []int{5, 5, 5, 10, 20}
	res := lemonadeChange(bills)
	fmt.Println(res)
}
