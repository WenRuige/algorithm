package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits); {
		if bits[i] == 1{
			i = i +2
		}else {
			if i == len(bits)-1{
				return true
			}
			i++
		}
	}
	return false

}

func main() {
	bits := []int{1,1,0}
	res := isOneBitCharacter(bits)
	fmt.Println(res)

}
