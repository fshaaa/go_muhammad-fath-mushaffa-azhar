package main

import "fmt"

func main() {
	var studentScore int = 80

	if studentScore <= 100 {
		if studentScore >= 80 {
			fmt.Println("A")
		} else if studentScore >= 65 {
			fmt.Println("B")
		} else if studentScore >= 50 {
			fmt.Println("C")
		} else if studentScore >= 35 {
			fmt.Println("D")
		} else if studentScore >= 0 {
			fmt.Println("E")
		} else {
			fmt.Println("Nilain Invalid")
		}
	} else {
		fmt.Println("Nilai Invalid")
	}
}
