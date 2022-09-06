package main

import "fmt"

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	item := make(map[string]int)
	for _, value := range items {
		item[value]++
	}

	result := make([]pair, 0)
	for key, value := range item {
		product := pair{key, value}
		result = append(result, product)
	}

	return result
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
