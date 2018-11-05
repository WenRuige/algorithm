package main

import (
	"fmt"
	"math"
)

//func constructRectangle(area int) []int {
//	temp := math.Sqrt(float64(area))
//	res := []int{}
//	if temp*temp == float64(area) {
//		return []int{int(temp), int(temp)}
//	}
//	temp = math.Floor(temp)
//	fmt.Println(temp)
//	second := area / int(temp)
//
//	if second > int(temp) {
//		res = append(res, second)
//		res = append(res, int(temp))
//	} else {
//		res = append(res, int(temp))
//		res = append(res, second)
//
//	}
//
//	return res
//}

func constructRectangle(area int) []int {
	temp := math.Sqrt(float64(area))
	res := []int{}
	if int(temp)*int(temp) == area {
		return []int{int(temp), int(temp)}
	}
	temp = math.Floor(temp)

	for i := temp; i > 0; i-- {
		fmt.Println(area % int(i))
		if area%int(i) == 0 {

			res = append(res, area/int(i))
			res = append(res, int(i))
			break
		}

	}
	return res
}

func main() {
	res := constructRectangle(2)
	fmt.Println(res)
}
