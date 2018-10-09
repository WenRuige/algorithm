package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	newArr := []int{}
	for i := 0; i < len(ops); i++ {
		// 表示最后一个回合是无效的,应该被移除
		if ops[i] == "C" {
			newArr = newArr[0 : len(newArr)-1]
			// 表示本轮获得的得分是前一轮有效 回合得分的两倍
		} else if ops[i] == "D" {
			newArr = append(newArr, newArr[len(newArr)-1]*2)
			//表示本轮获得的得分是前两轮有效 回合得分的总和
		} else if ops[i] == "+" {
			res := newArr[len(newArr)-1] + newArr[len(newArr)-2]
			newArr = append(newArr, res)
		} else {
			newNum, _ := strconv.Atoi(ops[i])
			newArr = append(newArr, newNum)
		}
	}
	sum := 0
	for _, v := range newArr {
		sum = sum + v
	}
	return sum

}

func main() {

	str := []string{"5", "2", "C", "D", "+"}
	res := calPoints(str)
	fmt.Println(res)
}
