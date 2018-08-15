package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	sum := []int{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				sum = append(sum, nums1[i])
				break
			}

		}
	}
	sum = RemoveRepByLoop(sum)

	return sum
}

func RemoveRepByLoop(slc []int) []int {
	result := []int{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 2, 2, 1, 3}
	nums2 := []int{2, 2, 2, 2, 3}
	res := intersection(nums1, nums2)
	fmt.Println(res)
}
