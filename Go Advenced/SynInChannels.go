package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)

	go func() {
		fmt.Println("Goroutine start...")
		ch <- "A"
		ch <- "B"
		ch <- "C"
		close(ch)
	}()

	for x := range ch {
		fmt.Println("Value:", x)
	}
}