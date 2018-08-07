package main

func reverse(x int) int {

	res := 0

	for x != 0 {
		pop := x % 10
		x = x / 10
		res = res*10 + pop
	}
	if res > 1<<31 || -res > 1<<31 {
		return 0
	}
	return res
}
