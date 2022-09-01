package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	var arrayC []string
	same := []int{}
	arrayC = append(arrayA, arrayB...)
	//fmt.Println("c = ", arrayC)
	for idx, value := range arrayC {
		for idx2 := idx + 1; idx2 < len(arrayC); idx2++ {
			if value == arrayC[idx2] {
				same = append(same, idx2)
				break
			}
		}
	}
	for _, value := range same {
		//len1 := len(arrayC)
		arrayC = append(arrayC[:value], arrayC[value+1:]...)
	}
	return arrayC
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimatsu"}, []string{"devil jin", "yoshimatsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{""}, []string{}))
}
