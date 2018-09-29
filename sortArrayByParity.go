package main

import "fmt"

func sortArrayByParity(A []int) []int {
	oddnumber := []int{} // 奇数
	evennumber := []int{}
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			evennumber = append(evennumber, A[i])
		} else {
			oddnumber = append(oddnumber, A[i])
		}
	}

	evennumber = append(evennumber, oddnumber...)
	return evennumber
}

func main() {
	nums := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(nums))
}
