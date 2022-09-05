package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	result := []int{}
	max := 0
	for _, card := range cards {
		tempMax := 0
		check := false
		for _, value := range card {
			tempMax += value
			for _, target := range deck {
				if value == target {
					check = true
				}
			}
			if tempMax > max && check {
				max = tempMax
				result = card
			}
		}
	}
	if max == 0 {
		return "tutup kartu"
	} else {
		return result
	}
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}
