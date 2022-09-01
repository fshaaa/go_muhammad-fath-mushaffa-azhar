package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var arr []int
	for _, value := range angka {
		num, _ := strconv.Atoi(string(value))
		arr = append(arr, num)
	}
	key := make(map[int]int)
	ans := []int{}

	for _, value := range arr {
		if n, _ := key[value]; n >= 0 {
			key[value] += 1
		}
	}
	for idx, value := range key {
		if value == 1 {
			ans = append(ans, idx)
		}
	}
	return ans
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76532752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
