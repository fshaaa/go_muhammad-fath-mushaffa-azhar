package main

import "fmt"

func Addition(num1, num2 int) int {
	return num1 + num2
}

func Subtraction(num1, num2 int) int {
	return num1 - num2
}

func Division(num1, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return 0
	}
	return num1 / num2
}

func Multiplication(num1, num2 int) int {
	return num1 * num2
}

func main() {
	fmt.Println(Addition(3, 7))
	fmt.Println(Subtraction(5, 2))
	fmt.Println(Division(4, 2))
	fmt.Println(Multiplication(2, 4))
}
