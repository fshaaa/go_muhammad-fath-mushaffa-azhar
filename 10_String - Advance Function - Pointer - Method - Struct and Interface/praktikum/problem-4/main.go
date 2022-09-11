package main

import "fmt"

func getMinMax(numbers ...*int) (min, max int) {
	min = 100000000
	max = -1
	for _, value := range numbers {
		if *value > max {
			max = *value
		}
		if *value < min {
			min = *value
		}
	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min", min)
	fmt.Println("Nilai max", max)
}
