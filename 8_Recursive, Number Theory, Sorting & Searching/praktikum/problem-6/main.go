package main

import "fmt"

func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}

func MaximumBuyProduct(money int, productPrice []int) {
	insertionSort(productPrice)
	var buy int = 0
	for _, value := range productPrice {
		if money >= value {
			money -= value
			buy++
		} else {
			break
		}
	}
	fmt.Println(buy)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
	MaximumBuyProduct(10000, []int{1000, 3000, 1000, 2000, 10000})
	MaximumBuyProduct(4000, []int{7500, 3000, 25000, 2000})
	MaximumBuyProduct(0, []int{10000, 30000})
}
