package main

import (
	"fmt"
	"strconv"
)

func FindMinAndMax(arr []int) string {
	var min, minIdx int = 100, 0
	var max, maxIdx int = -100, 0
	for i, value := range arr {
		if value < min {
			min = value
			minIdx = i
		}
		if value > max {
			max = value
			maxIdx = i
		}
	}
	return ("min: " + strconv.Itoa(min) + " index: " + strconv.Itoa(minIdx) + " max: " +
		strconv.Itoa(max) + " index: " + strconv.Itoa(maxIdx))
}

func main() {
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
}
