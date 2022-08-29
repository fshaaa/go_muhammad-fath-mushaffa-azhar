package main

import "fmt"

const phi = 3.14

func calculateSurfaceArea(t float64, r float64) float64 {
	Lp := 2 * phi * r * (r + t)
	return Lp
}

func main() {
	var t, r float64
	t = 20.0
	r = 4.0
	fmt.Println(calculateSurfaceArea(t, r))

	fmt.Println("Masukan tinggi dan jari-jari:")
	fmt.Scanf("%f %f", &t, &r)
	fmt.Println(calculateSurfaceArea(t, r))
}
