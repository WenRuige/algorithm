package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(len(nums1) / 2)
	if len(nums1)%2 != 0 {
		return float64(nums1[len(nums1)/2])
	} else {
		return (float64(nums1[len(nums1)/2]) + float64(nums1[(len(nums1)/2)-1])) / 2
	}
}

func main() {
	result := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(result)
}
