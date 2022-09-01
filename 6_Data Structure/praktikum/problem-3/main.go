package main

import "fmt"

func PairSum(arr []int, target int) []int {
	ans := []int{}
	i := 0
	j := len(arr) - 1
	for {
		if arr[i]+arr[j] == target {
			ans = append(ans, i, j)
			break
		} else if arr[i]+arr[j] < target {
			i++
		} else if arr[i]+arr[j] > target {
			j--
		}
		if i >= j {
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}
