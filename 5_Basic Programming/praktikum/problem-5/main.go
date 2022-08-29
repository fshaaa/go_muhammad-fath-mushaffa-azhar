package main

import "fmt"

func palindrome(input string) bool {
	for i := 0; i < len(input); i++ {
		for j := len(input) - i - 1; j >= 0; {
			//fmt.Printf("%c == %c", input[i], input[j])
			if input[i] == input[j] {
				break
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("mistar"))
	fmt.Println(palindrome("lion"))
}
