package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	ans := []string{}
	key := make(map[string]bool)

	for _, value := range arrayA {
		if _, num := key[value]; !num {
			key[value] = true
			ans = append(ans, value)
		}
	}
	for _, value := range arrayB {
		if _, num := key[value]; !num {
			key[value] = true
			ans = append(ans, value)
		}
	}

	return ans
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimatsu"}, []string{"devil jin", "yoshimatsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{""}, []string{}))
}
