package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	//password := []string{,, , , , ,, , , , , , ,, , , , , , , , , , , }

	password := map[string]string{}

	password["a"] = ".-"
	password["b"] = "-..."
	password["c"] = "-.-."
	password["d"] = "-.."
	password["e"] = "."
	password["f"] = "..-."
	password["g"] = "--."
	password["h"] = "...."
	password["i"] = ".."
	password["j"] = ".---"
	password["k"] = "-.-"
	password["l"] = ".-.."
	password["m"] = "--"
	password["n"] = "-."
	password["o"] = "---"
	password["p"] = ".--."
	password["q"] = "--.-"
	password["r"] = ".-."
	password["s"] = "..."
	password["t"] = "-"
	password["u"] = "..-"
	password["v"] = "...-"
	password["w"] = ".--"
	password["x"] = "-..-"
	password["y"] = "-.--"
	password["z"] = "--.."

	str := []string{}
	for i := 0; i < len(words); i++ {
		temp := ""
		for j := 0; j < len(words[i]); j++ {
			//fmt.Println(string(words[i][j]))
			data := password[string(words[i][j])]
			temp = temp + data
		}
		str = append(str, temp)
	}

	// 建立映射表

	reflect := make(map[string]string, len(words))

	for i := 0; i < len(str); i++ {
		res := reflect[str[i]]
		if res == "" {
			reflect[str[i]] = "cool"
		}
	}

	return len(reflect)

}

func main() {

	words := []string{"gin", "zen", "gig", "msg"}
	res := uniqueMorseRepresentations(words)
	fmt.Println(res)
}
