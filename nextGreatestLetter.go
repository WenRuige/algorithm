package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	min := -1 << 32
	result := byte(0)
	for i := 0; i < len(letters); i++ {
		temp := int((target - letters[i]))
		if min < temp {
			min = temp
			result = letters[i]
			//fmt.Println(result)
		}
		//fmt.Println(target - letters[i])
	}
	return result

}

func main() {
	nums := []byte{'c', 'f', 'j'}
	res := nextGreatestLetter(nums, 'a')
	fmt.Println(res)
}
