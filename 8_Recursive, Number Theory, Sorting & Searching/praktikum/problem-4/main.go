package main

import "fmt"

func maxSequence(arr []int) int {
	var tempMax, max int = 0, -100
	for _, value := range arr {
		tempMax = tempMax + value
		if tempMax > max {
			max = tempMax
		}
		if tempMax < 0 {
			tempMax = 0
		}
	}
	return max
}

func main() {
	fmt.Println(maxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(maxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(maxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(maxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
