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

	result := []pair{name: name, count: count}
	for key, value := range item {
		result.name = key
		result.count = value
	}

	return result
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
