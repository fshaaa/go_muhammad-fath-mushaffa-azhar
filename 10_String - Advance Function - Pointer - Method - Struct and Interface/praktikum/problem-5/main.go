package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var avg float64
	for _, value := range s.score {
		avg += float64(value)
	}
	avg /= 5
	return avg
}

func (s Student) Min() (min int, name string) {
	min = 101
	for i, value := range s.score {
		if min > value {
			min = value
			name = s.name[i]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max = -1
	for i, value := range s.score {
		if max < value {
			max = value
			name = s.name[i]
		}
	}
	return max, name
}

func main() {
	var a = Student{}

	for i := 1; i < 6; i++ {
		var name string
		fmt.Print("Input ", i, " Student's name: ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input ", i, " Score: ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAvarage Score Students is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
