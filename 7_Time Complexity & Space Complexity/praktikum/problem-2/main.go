package main

import "fmt"

func pow(x, n int) int {
	arr := make(map[int]int)
	arr[0] = 1
	arr[1] = x
	for i:=2; i<=n; i++{
		if arr[i] == 0 {
			arr[i] = x * arr[i-1] 
		}
	}
	return arr[n]
}

func main(){
	fmt.Println(pow(2,3))
	fmt.Println(pow(5,3))
	fmt.Println(pow(10,2))
	fmt.Println(pow(2,5))
	fmt.Println(pow(7,3))
}