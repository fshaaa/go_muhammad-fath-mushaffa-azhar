package main

import (
	"fmt"
	"math"
)

func checkPrime(number int) bool {
	if number == 2 {
		return true
	}

	if number < 2 || number%2 == 0 {
		return false
	}

	for i := 3; float64(i) <= math.Sqrt(float64(number)); i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func primeX(number int) int {
	ans := 0
	count := 0
	for i := 1; count <= number; i++ {
		if checkPrime(i) {
			count++
		}
		if count == number {
			ans = i
			return ans
		}
	}
	return -1
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
