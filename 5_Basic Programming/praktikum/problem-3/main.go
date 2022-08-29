package main

import "fmt"

func factorialBilangan(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	var bilangan int
	fmt.Scanf("%d", &bilangan)
	factorialBilangan(bilangan)
}
