package main

import "fmt"

type NumArray struct {
	Val []int
}

func ConstructorT(nums []int) NumArray {
	res := new(NumArray)
	res.Val = nums
	return *res

}

func (this *NumArray) SumRange(i int, j int) int {
	total := 0
	for begin := i; begin <= j; begin++ {
		total = total + this.Val[begin]
	}
	return total
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := ConstructorT(nums)
	res := obj.SumRange(0, 2)
	fmt.Println(res)
}
