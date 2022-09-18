package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	countString := make(map[string]int)
	var mx sync.Mutex

	for _, value := range text {
		countString[string(value)]++
		time.Sleep(1 * time.Nanosecond)

		go func() {
			mx.Lock()
			fmt.Printf("%s: %d\n", string(value), countString[string(value)])
			mx.Unlock()
		}()
	}
}
