package main

import "fmt"

func playWithAsterisk(n int) {
	for i := 1; i <= 5; i++ {
		for j := n - i; j >= 0; j-- {
			fmt.Printf(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}

func main() {
	playWithAsterisk(5)
}
