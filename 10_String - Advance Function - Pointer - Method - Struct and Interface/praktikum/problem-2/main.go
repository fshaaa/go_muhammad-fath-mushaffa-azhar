package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	offset = offset % 26
	ans := ""
	for _, value := range input {
		//fmt.Println(value, offset)
		value = value + rune(offset)
		if value > 'z' {
			value = value - rune(26)
		}
		ans += string(value)
	}
	return ans
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
