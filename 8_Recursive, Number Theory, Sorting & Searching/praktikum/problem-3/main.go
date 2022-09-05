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

func primaSegiEmpat(high, wide, start int) {
	start++
	var total int
	for i := 1; i <= wide; i++ {
		for j := 1; j <= high; j++ {
			check := false
			for ; !check; start++ {
				if checkPrime(start) {
					total += start
					fmt.Printf("%d ", start)
					check = true
				}
			}
		}
		fmt.Println()
	}
	fmt.Println(total)
}

func main() {
	primaSegiEmpat(2, 3, 13)
	primaSegiEmpat(5, 2, 1)
}
