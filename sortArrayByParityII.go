package main

import "fmt"

func sortArrayByParityII(A []int) []int {
	jishu := []int{}
	oushu := []int{}
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			oushu = append(oushu, A[i])
		} else {
			jishu = append(jishu, A[i])
		}
	}
	newArr := []int{}
	for i := 0; i < len(A); i++ {
		if i%2 == 0 {
			newArr = append(newArr, oushu[len(oushu)-1])
			oushu = oushu[:len(oushu)-1]
		} else {
			newArr = append(newArr, jishu[len(jishu)-1])
			jishu = jishu[:len(jishu)-1]
		}
	}

	return newArr
}

func main() {
	nums := []int{4, 2, 5, 7}
	res := sortArrayByParityII(nums)
	fmt.Println(res)
}
