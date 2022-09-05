package main

import "fmt"

func fibonacci(number int) int {
	var memo = make([]int, 100)
	if number == 0 {
		return memo[number]
	}
	if number == 1 {
		memo[number] = 1
	}
	if memo[number] > 0 {
		return memo[number]
	} else {
		memo[number] = fibonacci(number-1) + fibonacci(number-2)
	}
	return memo[number]
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(12))
}
