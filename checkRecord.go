package main

func checkRecord(s string) bool {
	flagA := 0
	temp := 0
	max := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "A" {
			flagA++
		}

		if string(s[i]) == "L" {
			temp++
			if max < temp {
				max = temp
			}

		} else {
			temp = 0
		}
	}

	if flagA > 2 {
		return false
	}
	if flagA < 2 && max < 3 {
		return true
	}

	return false
}

func main() {
	checkRecord("LLLAL")
}
