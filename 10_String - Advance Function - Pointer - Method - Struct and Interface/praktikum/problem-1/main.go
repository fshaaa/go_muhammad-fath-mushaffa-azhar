package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if len(a) < len(b) {
		temp := a
		a = b
		b = temp
	}
	if strings.Contains(a, b) {
		return b
	}
	return "not same"
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
